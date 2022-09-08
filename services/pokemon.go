package services

import (
	"errors"
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

var ErrEvolve = errors.New("could not evolve")

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
	convertOptions := convert.DefaultOptions
	return convert.NewImageConverter().Image2ASCIIString(img, &convertOptions), nil
}

func FetchRandomPokemon() (*models.PokemonModel, error) {
	randomPokemonId := rand.Int63n(NumPokemon)
	isShiny := rand.Intn(OddsShiny) == 1
	pokemonIdString := strconv.FormatInt(randomPokemonId, 10)
	resPokemon, err := pokeapi.Pokemon(pokemonIdString);
	if err != nil {
		return nil, err
	}
	evolvesTo, err := getEvolvesTo(pokemonIdString)
	if err != nil {
		return nil, err
	}
	pkmn := models.NewPokemonModelFromFetch(resPokemon, evolvesTo, isShiny)
	return pkmn, nil
}

func getEvolvesTo(pkmnId string) (string, error) {
	species, err := pokeapi.PokemonSpecies(pkmnId);
	currStageName := species.Name
	if err != nil {
		return "", err
	}
	evolutionChainUrl := species.EvolutionChain.URL
	var i int
	for i = len(evolutionChainUrl)-2; evolutionChainUrl[i] != '/'; i-- {}
	evolutionChainId := evolutionChainUrl[i+1:len(evolutionChainUrl)-1]
	evolutionChain, err := pokeapi.EvolutionChain(evolutionChainId)
	if err != nil {
		return "", err
	}

	if len(evolutionChain.Chain.EvolvesTo) == 0 {
		return "", nil
	}
	if evolutionChain.Chain.Species.Name == currStageName {
		return evolutionChain.Chain.EvolvesTo[0].Species.Name, nil
	}
	
	if len(evolutionChain.Chain.EvolvesTo[0].EvolvesTo) == 0 {
		return "", nil
	}
	if evolutionChain.Chain.EvolvesTo[0].Species.Name == currStageName {
		return evolutionChain.Chain.EvolvesTo[0].EvolvesTo[0].Species.Name, nil
	}

	if evolutionChain.Chain.EvolvesTo[0].EvolvesTo[0].Species.Name == currStageName {
		return "", nil
	}
	
	return "", ErrEvolve
}

func FetchPokemon(id string, isShiny bool) (*models.PokemonModel, error) {
	resPokemon, err := pokeapi.Pokemon(id);
	if err != nil {
		return nil, err
	}
	evolvesTo, err := getEvolvesTo(id)
	if err != nil {
		return nil, err
	}
	pkmn := models.NewPokemonModelFromFetch(resPokemon, evolvesTo, isShiny)
	return pkmn, nil
}