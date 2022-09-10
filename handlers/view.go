package handlers

import (
	"strings"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
)

func PrintAllPokemon(allPkmn *[]models.PokemonModel) {
	if len(*allPkmn) == 0 {
		helpers.PrintNoPokemon()
		return
	}

	for _, pkmn := range *allPkmn {
		helpers.PrintView(&pkmn)
	}
}

func PrintPartyPokemon(allPkmn *[]models.PokemonModel) {
	partyCount := 0
	for _, pkmn := range *allPkmn {
		if pkmn.IsInParty {
			helpers.PrintView(&pkmn)
			partyCount++
		}
	}

	if partyCount == 0 {
		helpers.PrintEmptyParty()
	}
}

func PrintFilteredPokemon(allPkmn *[]models.PokemonModel, filter string) {
	lowerFilter := strings.ToLower(filter)
	filterCount := 0
	for _, pkmn := range *allPkmn {
			if(strings.Contains(strings.ToLower(pkmn.Name), lowerFilter) || strings.Contains(strings.ToLower(pkmn.Nickname), lowerFilter)) {
				helpers.PrintView(&pkmn)
				filterCount++
			}
	}
	if filterCount == 0 {
		helpers.PrintBadFilter()
	}
}