package main

import (
	"fmt"
	"log"

	"github.com/doruo/falloutdle/external/wiki"
	"github.com/doruo/falloutdle/internal/database"
)

func main() {

	// Database connection creation
	db := database.NewDatabaseConnection()

	client := wiki.NewWikiClient()

	// Create and select example
	/*
		char := character.NewCharacter("Vault Dweller", "Vault_Dweller").
			SetRace("Human").
			SetGender("Male").
			SetStatus("Alive").
			SetMainGame("Fallout").
			AddGame("Fallout").
			AddGame("Fallout 2").
			AddAffiliation("Vault 13")
	*/
	name := "Roger_Maxson"
	char, err := client.FetchCharacterByName(name)

	fmt.Println("/--- CHARACTER ---/")
	fmt.Println(char.String())

	if err != nil {
		log.Printf("Error while fetching character %s: %v", name, err)
	}

	db.Create(&char)
	db.First(&char, 1)
}
