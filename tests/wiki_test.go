package tests

import (
	"fmt"
	"testing"

	"github.com/doruo/falloutdle/src/infrastructure/wiki"
)

var wiki_client = wiki.NewWikiClient()
var character_name = "Roger_Maxson"
var content_show_length = 50 // Set value > 0 to display response content

func TestMediaWikiClient_GetPageContent(t *testing.T) {

	content, err := wiki_client.GetPageContent(character_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("%s...\n", content[:content_show_length])
	} else {
		fmt.Printf("No content found")
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
		fmt.Println(character.String())
	} else {
		fmt.Printf("No content found")
	}
}
