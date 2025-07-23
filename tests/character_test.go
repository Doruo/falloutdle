package tests

import (
	"testing"

	"github.com/doruo/falloutdle/internal/character"
)

func TestNewCharacter(t *testing.T) {
	char := character.NewCharacter("Dogmeat", "Dogmeat_FO4")

	if char.Name != "Dogmeat" {
		t.Errorf("expected Dogmeat, got %s", char.Name)
	}
}
