package wiki

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

// CategoryResponse represents the response when querying category members
type CategoryResponse struct {
	Query struct {
		CategoryMembers []struct {
			PageID int    `json:"pageid"`
			Title  string `json:"title"`
		} `json:"categorymembers"`
	} `json:"query"`
}

// /----- FUNCTIONS -----/

// NewWikiClient creates a new WikiClient instance
func NewWikiClient() *WikiClient {
	return &WikiClient{
		baseURL: "https://fallout.fandom.com/api.php",
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

	// Extract content from response
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

// GetCharacterPages retrieves character pages from specific categories
func (w *WikiClient) GetCharacterPages(categories []string) ([]string, error) {
	var allPages []string

	for _, category := range categories {
		pages, err := w.getCategoryMembers(category)
		if err != nil {
			return nil, fmt.Errorf("failed to get category members for %s: %w", category, err)
		}
		allPages = append(allPages, pages...)
	}

	// Remove duplicates
	unique := make(map[string]bool)
	var result []string
	for _, page := range allPages {
		if !unique[page] {
			unique[page] = true
			result = append(result, page)
		}
	}

	return result, nil
}

// getCategoryMembers retrieves all pages from a specific category
func (w *WikiClient) getCategoryMembers(category string) ([]string, error) {
	params := url.Values{}
	params.Add("action", "query")
	params.Add("list", "categorymembers")
	params.Add("cmtitle", "Category:"+category)
	params.Add("cmlimit", "500")
	params.Add("format", "json")

	fullURL := w.baseURL + "?" + params.Encode()

	resp, err := w.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var catResp CategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&catResp); err != nil {
		return nil, fmt.Errorf("json decode failed: %w", err)
	}

	var pages []string
	for _, member := range catResp.Query.CategoryMembers {
		pages = append(pages, member.Title)
	}

	return pages, nil
}

// /--- PARSE FUNCTIONS ---/

// ParseCharacterFromContent parses MediaWiki content to extract character information
func (w *WikiClient) ParseCharacterFromContent(title, content string) (*models.Character, error) {
	// Extract infobox content using regex
	infoboxRegex := regexp.MustCompile(`\{\{Infobox character(.*?)\}\}`)
	matches := infoboxRegex.FindStringSubmatch(content)

	if len(matches) < 2 {
		return nil, fmt.Errorf("no character infobox found in %s", title)
	}

	infoboxContent := matches[1]

	// Parse infobox fields
	character := &models.Character{
		WikiTitle: title,
		Name:      w.extractField(infoboxContent, "name"),
		Race:      w.cleanRaceField(w.extractField(infoboxContent, "race")),
		Gender:    w.extractField(infoboxContent, "gender"),
		Status:    w.extractField(infoboxContent, "status"),
		Role:      w.extractField(infoboxContent, "role"),
		ImageURL:  w.extractField(infoboxContent, "image"),
	}

	// If no name field, use title
	if character.Name == "" {
		character.Name = title
	}

	// Parse games and mentions
	character.Games = w.parseGamesList(w.extractField(infoboxContent, "games"))
	character.Mentions = w.parseGamesList(w.extractField(infoboxContent, "mentions"))

	// Set main game
	if len(character.Games) > 0 {
		character.MainGame = character.Games[0]
	}

	// Parse affiliation
	character.Affiliation = w.parseAffiliationList(w.extractField(infoboxContent, "affiliation"))

	// Parse titles
	character.Titles = w.parseTitlesList(w.extractField(infoboxContent, "titles"))

	// Set timestamps
	now := time.Now()
	character.CreatedAt = now
	character.UpdatedAt = now

	return character, nil
}

// extractField extracts a field value from infobox content
func (w *WikiClient) extractField(content, field string) string {
	// Pattern to match field = value
	pattern := fmt.Sprintf(`\|%s\s*=\s*([^|]+)`, field)
	regex := regexp.MustCompile(pattern)
	matches := regex.FindStringSubmatch(content)

	if len(matches) < 2 {
		return ""
	}

	// Clean the value
	value := strings.TrimSpace(matches[1])
	value = w.cleanWikiText(value)

	return value
}

// cleanWikiText removes wiki markup from text
func (w *WikiClient) cleanWikiText(text string) string {
	// Remove wiki links [[text]] or [[link|text]]
	linkRegex := regexp.MustCompile(`\[\[([^|\]]+)(?:\|([^\]]+))?\]\]`)
	text = linkRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := linkRegex.FindStringSubmatch(match)
		if len(parts) > 2 && parts[2] != "" {
			return parts[2] // Use display text
		}
		return parts[1] // Use link text
	})

	// Remove references {{ref}}
	refRegex := regexp.MustCompile(`\{\{[^}]+\}\}`)
	text = refRegex.ReplaceAllString(text, "")

	// Remove HTML tags
	htmlRegex := regexp.MustCompile(`<[^>]+>`)
	text = htmlRegex.ReplaceAllString(text, "")

	// Remove ref tags
	refTagRegex := regexp.MustCompile(`<ref[^>]*>.*?</ref>`)
	text = refTagRegex.ReplaceAllString(text, "")

	// Clean up extra whitespace
	text = strings.ReplaceAll(text, "\n", " ")
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

// parseGamesList parses a comma-separated list of game codes
func (w *WikiClient) parseGamesList(gamesStr string) []string {
	if gamesStr == "" {
		return []string{}
	}

	games := strings.Split(gamesStr, ",")
	var result []string

	for _, game := range games {
		game = strings.TrimSpace(game)
		if game != "" {
			result = append(result, game)
		}
	}

	return result
}

// parseAffiliationList parses affiliation field which can be multi-line
func (w *WikiClient) parseAffiliationList(affiliationStr string) []string {
	if affiliationStr == "" {
		return []string{}
	}

	// Split by * (bullet points) and clean
	affiliations := strings.Split(affiliationStr, "*")
	var result []string

	for _, affiliation := range affiliations {
		affiliation = strings.TrimSpace(affiliation)
		if affiliation != "" {
			result = append(result, affiliation)
		}
	}

	return result
}

// parseTitlesList parses titles field
func (w *WikiClient) parseTitlesList(titlesStr string) []string {
	if titlesStr == "" {
		return []string{}
	}

	// Remove quotes and split by comma
	titlesStr = strings.ReplaceAll(titlesStr, "\"", "")
	titles := strings.Split(titlesStr, ",")
	var result []string

	for _, title := range titles {
		title = strings.TrimSpace(title)
		if title != "" {
			result = append(result, title)
		}
	}

	return result
}

// cleanRaceField specifically cleans race field which often has HTML breaks
func (w *WikiClient) cleanRaceField(race string) string {
	// Remove <br /> tags and take first race mentioned
	race = strings.ReplaceAll(race, "<br />", " ")
	race = strings.ReplaceAll(race, "<br>", " ")

	// If multiple races, take the first one
	if strings.Contains(race, " ") {
		parts := strings.Split(race, " ")
		race = parts[0]
	}

	return race
}

// GetDefaultCharacterCategories returns the default categories to scrape for characters
func (w *WikiClient) GetDefaultCharacterCategories() []string {
	return []string{
		"Fallout characters",
		"Fallout 2 characters",
		"Fallout 3 characters",
		"Fallout: New Vegas characters",
		"Fallout 4 characters",
		"Fallout 76 characters",
		"Fallout Tactics characters",
	}
}
