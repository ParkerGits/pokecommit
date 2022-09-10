package handlers

import (
	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
)

func HandleAddPokemon(allPkmn *[]models.PokemonModel) error {
	partyPkmn := helpers.FilterPkmn(allPkmn, func(pkmn models.PokemonModel) bool {
		return pkmn.IsInParty
	})
	boxPkmn := helpers.FilterPkmn(allPkmn, func(pkmn models.PokemonModel) bool {
		return !pkmn.IsInParty
	})
	
	if len(*boxPkmn) == 0 {
		helpers.PrintNoPokemonInBox()
		return nil
	}

	toAdd, err := selectPokemon(boxPkmn, "Select a Pokemon from your box to add to your party")
	if err != nil {
		return err
	}
	if toAdd == nil {
		return nil
	}
	if len(*partyPkmn) < 6 {
		if err := addPokemonToParty(toAdd); err != nil {
			return err
		}
		return nil
	}
	toReplace, err := selectPokemon(partyPkmn, "Your party is full! Select a Pokemon from your party to replace:")
	if err != nil {
		return err
	}
	return replacePokemonInParty(toReplace, toAdd)
}

func HandleDepositPokemon(allPkmn *[]models.PokemonModel) error {
	partyPkmn := helpers.FilterPkmn(allPkmn, func(pkmn models.PokemonModel) bool {
		return pkmn.IsInParty
	})
	if len(*partyPkmn) == 0 {
		helpers.PrintEmptyParty()
		return nil
	}
	if len(*partyPkmn) == 1 {
		helpers.PrintCannotDepositLastPokemon()
		return nil
	}

	toDeposit, err := selectPokemon(partyPkmn, "Select a Pokemon from your party to deposit")
	if err != nil {
		return err
	}
	if toDeposit == nil {
		return nil
	}
	return depositPokemonIntoBox(toDeposit)
}

func replacePokemonInParty(toReplace *models.PokemonModel, toAdd *models.PokemonModel) error {
	return models.ReplaceInParty(toReplace, toAdd)
}

func addPokemonToParty(pkmn *models.PokemonModel) error {
	pkmn.IsInParty = true
	if err := models.UpdatePokemon(pkmn); err != nil {
		return err
	}
	helpers.PrintAddedToParty(pkmn)
	return nil
}

func depositPokemonIntoBox(pkmn *models.PokemonModel) error {
	pkmn.IsInParty = false
	if err := models.UpdatePokemon(pkmn); err != nil {
		return err
	}
	helpers.PrintDeposited(pkmn)
	return nil
}