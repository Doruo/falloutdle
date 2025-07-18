package tests

import (
	"fmt"
	"testing"

	"github.com/doruo/falloutdle/src/infrastructure/wiki"
)

var wiki_client = wiki.NewWikiClient()
var character_name = "Roger_Maxson"
var content_show_length = 0 // Set value > 0 to display response content
func TestMediaWikiClient_GetPageContent(t *testing.T) {

	if _, err := wiki_client.GetPageContent(character_name); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestMediaWikiClient_ParseCharacterFromContent(t *testing.T) {

	content, err := wiki_client.GetPageContent(character_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("%s...\n", content[:content_show_length])
	} else {
		fmt.Printf("No content found")
	}

	character, err := wiki_client.ParseCharacterFromContent(character_name, content)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("CHARACTER: %v", character)
	} else {
		fmt.Printf("No content found")
	}
}
