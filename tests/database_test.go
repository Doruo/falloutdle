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

	db := database.NewDatabaseConnection()
	repository := character.NewCharacterRepository(db)
	client := wiki.NewWikiClient()
	title := "Roger_Maxson"

	char, err := client.FetchCharacterByName(title)
	if err != nil {
		log.Printf("Error while fetching character %s: %v", title, err)
	}

	fmt.Println("/--- CHARACTER ---/")

	fmt.Println(char.String())

	fmt.Println("/--- DELETE ---/")
	repository.DeleteByWikiTitle(title)

	fmt.Println("/--- ADD ---/")
	repository.Add(char)

	fmt.Println("/--- SELECT ---/")
	result, err := repository.GetByWikiTitle(title)
	if err != nil {
		log.Printf("Error while fetching character %s: %v", title, err)
	}

	fmt.Println(result.String())
}
