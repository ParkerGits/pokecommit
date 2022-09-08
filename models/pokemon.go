package models

import (
	"github.com/mtslzr/pokeapi-go/structs"
	"gorm.io/gorm"
)

type PokemonModel struct {
	gorm.Model
	PokeId uint16
	Name string
	Nickname string
	AsciiFrontSpriteUrl string
	AsciiBackSpriteUrl string
	EvolvesTo string
	IsShiny bool
	Type1 string
	Type2 string
	IsInParty bool
}

func NewPokemonModelFromFetch(fetched structs.Pokemon, evolvesTo string, isShiny bool) *PokemonModel {
	pkmn := new(PokemonModel)
	pkmn.PokeId = uint16(fetched.ID)
	pkmn.Name = fetched.Name
	pkmn.IsShiny = isShiny
	pkmn.EvolvesTo = evolvesTo
	pkmn.Type1 = fetched.Types[0].Type.Name
	if len(fetched.Types) > 1 {
		pkmn.Type2 = fetched.Types[1].Type.Name
	}
	if pkmn.IsShiny {
		pkmn.AsciiFrontSpriteUrl = fetched.Sprites.FrontShiny
		pkmn.AsciiBackSpriteUrl = fetched.Sprites.BackShiny
		return pkmn;
	}
	pkmn.AsciiFrontSpriteUrl = fetched.Sprites.FrontDefault
	pkmn.AsciiBackSpriteUrl = fetched.Sprites.BackDefault
	return pkmn;
}

func (p *PokemonModel) TableName() string {
	return "pokemon"
}

func (p *PokemonModel) EvolveInto(evolution *PokemonModel) error {
	p.PokeId = evolution.PokeId
	p.Name = evolution.Name
	p.AsciiFrontSpriteUrl = evolution.AsciiFrontSpriteUrl
	p.AsciiBackSpriteUrl = evolution.AsciiBackSpriteUrl
	p.EvolvesTo = evolution.EvolvesTo
	p.Type1 = evolution.Type1
	p.Type2 = evolution.Type2
	return UpdatePokemon(p)
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