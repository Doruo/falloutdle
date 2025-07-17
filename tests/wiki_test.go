package tests

import (
	"fmt"
	"testing"

	"github.com/doruo/falloutdle/src/infrastructure/wiki"
)

func TestMediaWikiClient_GetPage(t *testing.T) {

	fmt.Println("Testing getWikiPageContent...")

	client := wiki.NewWikiClient()
	page_name := "Roger_Maxson"
	content_show_lenght := 1990

	content, err := client.GetPageContent(page_name)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) > 0 {
		fmt.Printf("%s...\n", content[:content_show_lenght])
	} else {
		fmt.Printf("No content found")
	}
}
