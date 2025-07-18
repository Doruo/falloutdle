package wiki

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/doruo/falloutdle/src/domains/models"
)

// /----- STRUCTS -----/

// WikiClient handles communication with Fallout Wiki API
type WikiClient struct {
	baseURL    string
	httpClient *http.Client
}

// WikiResponse represents the MediaWiki API response structure
type WikiResponse struct {
	Query struct {
		Pages map[string]WikiPage `json:"pages"`
	} `json:"query"`
}

// WikiPage represents a single wiki page
type WikiPage struct {
	PageID    int            `json:"pageid"`
	Title     string         `json:"title"`
	Revisions []WikiRevision `json:"revisions"`
}

// WikiRevision represents a page revision
type WikiRevision struct {
	Slots struct {
		Main struct {
			ContentFormat string `json:"contentformat"`
			ContentModel  string `json:"contentmodel"`
			Content       string `json:"*"`
		} `json:"main"`
	} `json:"slots"`
}

// CategoryResponse represents the category members API response response when querying category members
type CategoryResponse struct {
	Query struct {
		CategoryMembers []CategoryMember `json:"categorymembers"`
	} `json:"query"`
}

// CategoryMember represents a member of a category
type CategoryMember struct {
	PageID int    `json:"pageid"`
	Title  string `json:"title"`
	NS     int    `json:"ns"`
}

// /----- FUNCTIONS -----/

// NewWikiClient creates a new WikiClient instance
func NewWikiClient() *WikiClient {
	return &WikiClient{
		baseURL: os.Getenv("WIKI_API_URL"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
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

	fullURL := w.baseURL + "?" + params.Encode()

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

// GetCategoryMembers retrieves all pages in a specific category
func (w *WikiClient) GetCategoryMembers(category string) ([]CategoryMember, error) {

	params := url.Values{}
	params.Add("action", "query")
	params.Add("list", "categorymembers")
	params.Add("cmtitle", "Category:"+category)
	params.Add("cmlimit", "500") // Maximum allowed
	params.Add("format", "json")

	fullURL := w.baseURL + "?" + params.Encode()

	resp, err := w.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var catResp CategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&catResp); err != nil {
		return nil, fmt.Errorf("JSON parsing failed: %w", err)
	}

	return catResp.Query.CategoryMembers, nil
}

// /--- PARSE FUNCTIONS ---/

// ParseCharacterFromContent parses MediaWiki content to extract character information
func (w *WikiClient) ParseCharacterFromContent(title, content string) (*models.Character, error) {
	// Find the infobox character section
	infoboxRegex := regexp.MustCompile(`(?s)\{\{Infobox character(.*?)\}\}`)
	matches := infoboxRegex.FindStringSubmatch(content)

	if len(matches) < 1 {
		return nil, fmt.Errorf("no character infobox found in page: %s", title)
	}

	infoboxContent := matches[1]

	character := &models.Character{
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
			character.Name = w.cleanWikiText(value)
		case "games":
			character.Games = models.NormalizeGameCodes(value)
		case "mentions":
			character.Mentions = models.NormalizeGameCodes(value)
		case "race":
			character.Race = w.cleanWikiText(value)
		case "gender":
			character.Gender = w.cleanWikiText(value)
		case "status":
			character.Status = w.cleanWikiText(value)
		case "affiliation":
			character.Affiliation = w.parseAffiliation(value)
		case "role":
			character.Role = w.cleanWikiText(value)
		case "titles":
			character.Titles = w.parseTitles(value)
		case "image":
			character.ImageURL = w.parseImageURL(value)
		}
	}

	// Set main game
	character.MainGame = character.GetMainGame()

	return character, nil
}

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

// GetCharactersByCategory retrieves characters from a specific category
func (w *WikiClient) GetCharactersByCategory(category string) ([]*models.Character, error) {
	members, err := w.GetCategoryMembers(category)
	if err != nil {
		return nil, fmt.Errorf("failed to get category members: %w", err)
	}

	var characters []*models.Character

	for _, member := range members {
		// Skip non-article pages (ns != 0)
		if member.NS != 0 {
			continue
		}

		content, err := w.GetPageContent(member.Title)
		if err != nil {
			fmt.Printf("Warning: failed to get content for %s: %v\n", member.Title, err)
			continue
		}

		character, err := w.ParseCharacterFromContent(member.Title, content)
		if err != nil {
			fmt.Printf("Warning: failed to parse character %s: %v\n", member.Title, err)
			continue
		}

		characters = append(characters, character)

		// Add small delay to be respectful to the API
		time.Sleep(100 * time.Millisecond)
	}

	return characters, nil
}
