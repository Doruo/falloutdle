package main

import (
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
)

func main() {

	// Database connection creation
	db := database.NewDatabaseConnection()

	// Create and select example
	char := character.NewCharacter("Vault Dweller", "Vault_Dweller").
		SetRace("Human").
		SetGender("Male").
		SetStatus("Alive").
		SetMainGame("Fallout").
		AddGame("Fallout").
		AddGame("Fallout 2").
		AddAffiliation("Vault 13")

	db.Create(&char)
	db.First(&char, 1)
}
