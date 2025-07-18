package tests

import (
	"testing"

	"github.com/doruo/falloutdle/src/domains/models/wiki"
)

var wiki_client = wiki.NewWikiClient()
var character_name = "Roger_Maxson"
var content_show_length = 50 // Set value > 0 to display response content

func TestMediaWikiClient_GetPageContent(t *testing.T) {

	if content, err := wiki_client.GetPageContent(character_name); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	} else if len(content) > 0 {
		t.Logf("%s...\n", content[:content_show_length])
	} else {
		t.Fatalf("No content found")
	}
}

func TestMediaWikiClient_ParseCharacterFromContent(t *testing.T) {

	content, err := wiki_client.GetPageContent(character_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	character, err := wiki_client.ParseCharacterFromContent(character_name, content)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		t.Log(character.StringCompact())
	} else {
		t.Fatalf("No content found")
	}
}

func TestMediaWikiClient_FetchAllCharacters(t *testing.T) {
	characters, err := wiki_client.FetchAllCharacters()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(characters) > 0 {
		for _, character := range characters {
			t.Log(character.Name)
		}
	} else {
		t.Fatalf("No characters found")
	}
}
