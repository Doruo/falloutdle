package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Structures pour parser la réponse de l'API MediaWiki
type WikiResponse struct {
	Query struct {
		Pages map[string]WikiPage `json:"pages"`
	} `json:"query"`
}

type WikiPage struct {
	PageID    int            `json:"pageid"`
	Title     string         `json:"title"`
	Revisions []WikiRevision `json:"revisions"`
}

type WikiRevision struct {
	Slots struct {
		Main struct {
			ContentFormat string `json:"contentformat"`
			ContentModel  string `json:"contentmodel"`
			Content       string `json:"*"`
		} `json:"main"`
	} `json:"slots"`
}

// Extract wiki page content from title
func getWikiPageContent(title string) (string, error) {
	baseURL := "https://fallout.fandom.com/api.php"

	// Construire les paramètres de l'URL
	params := url.Values{}
	params.Add("action", "query")
	params.Add("prop", "revisions")
	params.Add("rvprop", "content")
	params.Add("rvslots", "main")
	params.Add("format", "json")
	params.Add("titles", title)

	// complete url creation
	fullURL := baseURL + "?" + params.Encode()
	fmt.Println(fullURL)

	// HTTP request
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("erreur HTTP: %w", err)
	}
	defer resp.Body.Close()

	// parse JSON response
	var wikiResp WikiResponse
	if err := json.NewDecoder(resp.Body).Decode(&wikiResp); err != nil {
		return "", fmt.Errorf("erreur parsing JSON: %w", err)
	}

	// extract page content
	for _, page := range wikiResp.Query.Pages {
		if len(page.Revisions) > 0 {
			return page.Revisions[0].Slots.Main.Content, nil
		}
	}

	// error case
	return "", fmt.Errorf("contenu non trouvé pour: %s", title)
}

func main() {

	// test
	content, err := getWikiPageContent("Vault_Dweller")
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}

	// display test content result
	fmt.Printf("Contenu récupéré (%d caractères):\n", len(content))
	if len(content) > 500 {
		fmt.Printf("%s...\n", content[:500])
	} else {
		fmt.Printf("%s\n", content)
	}
}
