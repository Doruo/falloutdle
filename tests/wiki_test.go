package tests

import (
	"fmt"
	"testing"

	"github.com/doruo/falloutdle/src/infrastructure/wiki"
)

var wiki_client = wiki.NewWikiClient()
var page_name = "John_Hancock"
var content_show_length = 0

func TestMediaWikiClient_GetPageContent(t *testing.T) {

	if _, err := wiki_client.GetPageContent(page_name); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestMediaWikiClient_ParseCharacterFromContent(t *testing.T) {

	content, err := wiki_client.GetPageContent(page_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("%s...\n", content[:content_show_length])
	} else {
		fmt.Printf("No content found")
	}

	character, err := wiki_client.ParseCharacterFromContent(content)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("Character result: %v", character)
	} else {
		fmt.Printf("No content found")
	}
}
