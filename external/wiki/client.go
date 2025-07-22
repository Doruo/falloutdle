package wiki

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/doruo/falloutdle/internal/character"
)

// Fallout Fandom Wiki API URL
var wiki_api_url = "https://fallout.fandom.com/api.php"

// /----- STRUCTS -----/

// WikiClient handles communication with Fallout Wiki API
type WikiClient struct {
	baseURL    string
	httpClient *http.Client
}

// /----- FUNCTIONS -----/

// NewWikiClient creates a new WikiClient instance
// Timeout specifies a time limit in seconds for requests made by this Client.
func NewWikiClient() *WikiClient {
	return &WikiClient{
		baseURL: wiki_api_url,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// GetPageContent retrieves the raw content of a wiki page
func (w *WikiClient) GetPageContent(title string) (string, error) {

	params := url.Values{}
	params.Add("action", "query")
	params.Add("prop", "revisions")
	params.Add("rvprop", "content")
	params.Add("rvslots", "main")
	params.Add("format", "json")
	params.Add("titles", title)

	// Full wiki api request parse
	fullURL := w.baseURL + "?" + params.Encode()

	// HTTP request to wiki api
	resp, err := w.httpClient.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var wikiResp WikiResponse
	if err := json.NewDecoder(resp.Body).Decode(&wikiResp); err != nil {
		return "", fmt.Errorf("json decode failed: %w", err)
	}

	// Extract content from response first page
	for _, page := range wikiResp.Query.Pages {

		if page.PageID == -1 {
			return "", fmt.Errorf("page not found: %s", title)
		}
		if len(page.Revisions) > 0 {
			return page.Revisions[0].Slots.Main.Content, nil
		}
	}

	return "", fmt.Errorf("no content found for page: %s", title)
}

// /--- PARSE FUNCTIONS ---/

// ParseCharacterFromContent parses MediaWiki content to extract character information
func (w *WikiClient) ParseCharacterFromContent(title, content string) (*character.Character, error) {
	// Find the infobox character section
	infoboxRegex := regexp.MustCompile(`(?s)\{\{Infobox character(.*?)\}\}`)
	matches := infoboxRegex.FindStringSubmatch(content)

	if len(matches) < 1 {
		return nil, fmt.Errorf("no character infobox found in page: %s", title)
	}

	infoboxContent := matches[1]

	c := &character.Character{
		WikiTitle: title,
		Name:      title, // Default name, will be overridden if 'name' field exists
	}

	// Parse infobox fields
	lines := strings.Split(infoboxContent, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "|") {
			continue
		}

		// Remove the | prefix and split by =
		line = strings.TrimPrefix(line, "|")
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Skip empty values
		if value == "" {
			continue
		}

		// Parse different fields
		switch key {
		case "name":
			c.Name = w.cleanWikiText(value)
		case "games":
			c.Games = character.NormalizeGameCodes(value)
		case "mentions":
			c.Mentions = character.NormalizeGameCodes(value)
		case "race":
			c.Race = w.cleanWikiText(value)
		case "gender":
			c.Gender = w.cleanWikiText(value)
		case "status":
			c.Status = w.cleanWikiText(value)
		case "affiliation":
			c.Affiliation = w.parseAffiliation(value)
		case "role":
			c.Role = w.cleanWikiText(value)
		case "titles":
			c.Titles = w.parseTitles(value)
		case "image":
			c.ImageURL = w.parseImageURL(value)
		}
	}

	// Set main game
	c.MainGame = c.GetMainGame()

	return c, nil
}

// /---- FETCH FUNCTIONS -----/

func (w *WikiClient) FetchAllCharacters() (characters []*character.Character, e error) {

	for _, game_code := range character.AllGameCodes {
		temp_characters, err := w.FetchCharactersByGame(game_code)

		// Error case
		if err != nil {
			log.Printf("Error fetching characters for game %s: %v", game_code, err)
			return nil, e
		}

		// []*Character to []Character conversion before appending
		for _, character := range temp_characters {
			if character != nil {
				characters = append(characters, character)
			}
		}
	}
	return
}

func (w *WikiClient) FetchCharactersByGame(game character.GameCode) ([]*character.Character, error) {
	categoryTitle := fmt.Sprintf("Category:%s_characters", strings.ReplaceAll(game.GameFullName(), " ", "_"))

	params := url.Values{}
	params.Set("action", "query")
	params.Set("format", "json")
	params.Set("list", "categorymembers")
	params.Set("cmtitle", categoryTitle)
	params.Set("cmlimit", "500")

	var characters []*character.Character
	cmContinue := ""

	for {
		iterParams := url.Values{}
		for k, vs := range params {
			for _, v := range vs {
				iterParams.Add(k, v)
			}
		}
		if cmContinue != "" {
			iterParams.Set("cmcontinue", cmContinue)
		}

		fullURL := w.baseURL + "?" + iterParams.Encode()

		resp, err := w.httpClient.Get(fullURL)
		if err != nil {
			return nil, fmt.Errorf("HTTP request failed: %w", err)
		}
		defer resp.Body.Close()

		var result struct {
			Continue struct {
				CmContinue string `json:"cmcontinue"`
			} `json:"continue"`
			Query struct {
				CategoryMembers []struct {
					Title string `json:"title"`
					NS    int    `json:"ns"`
				} `json:"categorymembers"`
			} `json:"query"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}

		for _, member := range result.Query.CategoryMembers {
			if member.NS != 0 {
				continue
			}

			content, err := w.GetPageContent(member.Title)
			if err != nil {
				log.Printf("Failed to get content for %s: %v", member.Title, err)
				continue
			}

			character, err := w.ParseCharacterFromContent(member.Title, content)
			if err != nil {
				log.Printf("Failed to parse character %s: %v", member.Title, err)
				continue
			}

			characters = append(characters, character)
		}

		if result.Continue.CmContinue == "" {
			break
		}
		cmContinue = result.Continue.CmContinue
	}

	return characters, nil
}

// /---- UTILITY FUNCTIONS -----/

// cleanWikiText removes wiki markup from text
func (w *WikiClient) cleanWikiText(text string) string {
	// Remove wiki links [[Text]] or [[Link|Text]]
	linkRegex := regexp.MustCompile(`\[\[([^\]|]+)(\|[^\]]+)?\]\]`)
	text = linkRegex.ReplaceAllStringFunc(text, func(match string) string {
		// Extract the display text (after |) or the link text
		if strings.Contains(match, "|") {
			parts := strings.Split(match, "|")
			if len(parts) > 1 {
				return strings.TrimSuffix(parts[1], "]]")
			}
		}
		// Remove [[ and ]]
		return strings.TrimSuffix(strings.TrimPrefix(match, "[["), "]]")
	})

	// Remove HTML tags and <br> elements
	htmlRegex := regexp.MustCompile(`<[^>]+>`)
	text = htmlRegex.ReplaceAllString(text, "")

	// Remove references like {{cn}} or <ref>...</ref>
	refRegex := regexp.MustCompile(`\{\{[^}]+\}\}|<ref[^>]*>.*?</ref>`)
	text = refRegex.ReplaceAllString(text, "")

	// Clean up multiple spaces and newlines
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

// parseAffiliation parses affiliation field which can be a list
func (w *WikiClient) parseAffiliation(value string) []string {
	// Split by * or newlines for list items
	lines := strings.Split(value, "\n")
	var affiliations []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimSpace(line)

		if line != "" {
			clean := w.cleanWikiText(line)
			if clean != "" {
				affiliations = append(affiliations, clean)
			}
		}
	}

	// If no list format, try comma separation
	if len(affiliations) == 0 {
		parts := strings.Split(value, ",")
		for _, part := range parts {
			clean := w.cleanWikiText(strings.TrimSpace(part))
			if clean != "" {
				affiliations = append(affiliations, clean)
			}
		}
	}

	return affiliations
}

// parseTitles parses titles field
func (w *WikiClient) parseTitles(value string) []string {
	// Titles are often quoted strings
	titleRegex := regexp.MustCompile(`"([^"]+)"`)
	matches := titleRegex.FindAllStringSubmatch(value, -1)

	var titles []string
	for _, match := range matches {
		if len(match) > 1 {
			titles = append(titles, match[1])
		}
	}

	return titles
}

// parseImageURL extracts image filename (could be expanded to full URL)
func (w *WikiClient) parseImageURL(value string) string {
	// For now, just return the filename
	// Could be expanded to construct full Fandom image URLs
	return strings.TrimSpace(value)
}
