package helpers

import (
	"github.com/ParkerGits/pokecommit/services"
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
	PrintPkmn(pkmn, sprite)
	return nil
}

