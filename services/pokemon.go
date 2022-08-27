package services

import (
	"image/png"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ParkerGits/pokecommit/models"
	"github.com/mtslzr/pokeapi-go"
	"github.com/qeesung/image2ascii/convert"
)

const (
	NumPokemon = 898
	OddsShiny = 100
)

func FetchAsciiSprite(imgUrl string) (string, error) {
	resSprite, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer resSprite.Body.Close()
	img, err := png.Decode(resSprite.Body)
	if err != nil {
		return "", err
	}
	convertOptions := convert.DefaultOptions;
	convertOptions.StretchedScreen = true;
	return convert.NewImageConverter().Image2ASCIIString(img, &convert.DefaultOptions), nil
}

func FetchPokemon() (pkmn models.PokemonModel, err error){
	randomPokemonId := rand.Int63n(NumPokemon)
	resPokemon, err := pokeapi.Pokemon(strconv.FormatInt(randomPokemonId, 10));
	if err != nil {
		return pkmn, err
	}
	pkmn.Name = resPokemon.Name
	pkmn.IsShiny = rand.Intn(OddsShiny) == 1
	if pkmn.IsShiny {
		pkmn.AsciiSpriteUrl = resPokemon.Sprites.FrontShiny
		return pkmn, nil;
	}
	pkmn.AsciiSpriteUrl = resPokemon.Sprites.FrontDefault
	return pkmn, nil;
}