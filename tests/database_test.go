package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/doruo/falloutdle/external/wiki"
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
)

func TestAddCharacter(t *testing.T) {

	db := database.GetInstance()
	repository := character.NewCharacterRepository(db)
	client := wiki.NewWikiClient()
	title := "Roger_Maxson"

	char, error := client.FetchCharacterByName(title)
	if error != nil {
		log.Printf("Error while fetching character %s: %v", title, error)
	}

	fmt.Println("/--- CHARACTER ---/")

	fmt.Println(char.String())

	fmt.Println("/--- DELETE ---/")
	repository.DeleteByWikiTitle(title)

	fmt.Println("/--- ADD ---/")
	repository.Add(char)

	fmt.Println("/--- SELECT ---/")
	result, error := repository.GetByWikiTitle(title)
	if error != nil {
		log.Printf("Error while fetching character %s: %v", title, error)
	}

	fmt.Println(result.String())
}
