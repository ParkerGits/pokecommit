package handlers

import (
	"strings"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
)

func PrintAllPokemon(allPkmn *[]models.PokemonModel) {
	for _, pkmn := range *allPkmn {
		helpers.PrintView(&pkmn)
	}
}

func PrintPartyPokemon(allPkmn *[]models.PokemonModel) error {
	partyCount := 0
	for _, pkmn := range *allPkmn {
		if pkmn.IsInParty {
			helpers.PrintView(&pkmn)
		}
	}
	
	if partyCount == 0 {
		return ErrEmptyPrty
	}

	return nil
}

func PrintFilteredPokemon(allPkmn *[]models.PokemonModel, filter string) {
	lowerFilter := strings.ToLower(filter)
		for _, pkmn := range *allPkmn {
				if(strings.Contains(strings.ToLower(pkmn.Name), lowerFilter) || strings.Contains(strings.ToLower(pkmn.Nickname), lowerFilter)) {
					helpers.PrintView(&pkmn)
				}
		}
}