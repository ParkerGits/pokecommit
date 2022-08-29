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
	IsInParty bool
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

func GetParty(pokemon *[]PokemonModel) error {
	if result := db.Find(pokemon, PokemonModel{IsInParty: true}); result.Error != nil {
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

func UpdatePokemon(pokemon *PokemonModel) error {
	if result := db.Save(pokemon); result.Error != nil {
		return result.Error
	}
	return nil
}