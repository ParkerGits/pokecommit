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