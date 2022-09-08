package handlers

import (
	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
)

func PrintAllPokemon(allPkmn []models.PokemonModel) {
	for _, pkmn := range allPkmn {
		helpers.PrintView(&pkmn)
	}
}