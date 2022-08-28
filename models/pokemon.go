package models

import (
	"gorm.io/gorm"
)

type PokemonModel struct {
	gorm.Model
	PokeId uint16
	Name string
	Nickname string
	AsciiSpriteUrl string
	IsShiny bool
	Type1 string
	Type2 string
}

func (p *PokemonModel) TableName() string {
	return "pokemon"
}

func GetAllPokemon(pokemon *[]PokemonModel) error {
	if result := db.Find(pokemon); result.Error != nil {
		return result.Error
	}
	return nil
}

func CreatePokemon(pokemon *PokemonModel) error {
	if result := db.Create(pokemon); result.Error != nil {
		return result.Error
	}
	return nil
}