package tests

import (
	"testing"

	"github.com/doruo/falloutdle/external/wiki"
)

var client = wiki.NewWikiClient()
var character_name = "Roger_Maxson"
var content_show_length = 50 // Set value > 0 to display response content

func TestMediaWikiClient_GetPageContent(t *testing.T) {

	if content, err := client.FetchPageContent(character_name); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	} else if len(content) > 0 {
		t.Logf("%s...\n", content[:content_show_length])
	} else {
		t.Fatalf("No content found")
	}
}

func TestMediaWikiClient_ParseCharacterFromContent(t *testing.T) {

	content, err := client.FetchPageContent(character_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	character, err := client.ParseCharacterFromContent(character_name, content)

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

	t.Skip("Skipping expensive test") // REMOVE TO DO TEST

	characters, err := client.FetchAllCharacters()

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
