package helpers

import (
	"github.com/ParkerGits/pokecommit/services"
	"github.com/ParkerGits/pokecommit/models"
)

func CatchRandomPokemon() error {
	pkmn, err := services.FetchRandomPokemon()
	if err != nil {
		return err
	}
	sprite, err := services.FetchAsciiSprite(pkmn.AsciiSpriteUrl)
	if err != nil {
		return err
	}
	PrintEncounter(pkmn, sprite)
	if err := models.CreatePokemon(&pkmn); err != nil {
		return err
	}
	return nil
}