package helpers

import (
	"github.com/ParkerGits/pokecommit/models"
)

func MapToString(pkmn []models.PokemonModel, f func(models.PokemonModel, int) string) []string {
	mapped := make([]string, len(pkmn))
	for i, v := range pkmn {
			mapped[i] = f(v, i)
	}
	return mapped
}

func FilterPkmn(allPkmn *[]models.PokemonModel, callback func(models.PokemonModel) bool) *[]models.PokemonModel {
	partyPkmn := []models.PokemonModel{}
	for _, pkmn := range *allPkmn {
		if callback(pkmn){
			partyPkmn = append(partyPkmn, pkmn)
		}
	}
	return &partyPkmn
}