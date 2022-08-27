package main

// todo remove jpeg?

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ParkerGits/pokegit/models"
	"github.com/ParkerGits/pokegit/services"
)


type Sprites struct {
	Front []byte `json:"front_default"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	randPkmn, err := services.FetchPokemon()
	if err != nil {
		log.Fatal(err)
	}
	models.CreatePokemon(&randPkmn)
	allPkmn := []models.PokemonModel{}
	err = models.GetAllPokemon(&allPkmn)
	if err != nil {
		log.Fatal(err)
	}
	for _, pkmn := range allPkmn {
		fmt.Println(pkmn.Name)
		asciiSprite, err := services.FetchAsciiSprite(pkmn.AsciiSpriteUrl)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(asciiSprite)
	}
}