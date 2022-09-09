package handlers

import (
	"errors"
	"math/rand"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
	"github.com/ParkerGits/pokecommit/services"
	"github.com/manifoldco/promptui"
)

var (
	encounterActions = []string{
		"Catch",
		"Fight",
		"Run",
	}
	yesNo = []string{
		"Yes",
		"No",
	}
	ErrEmptyPrty = errors.New("empty party")
)

func EngageRandomEncounter() error {
	pkmn, err := services.FetchRandomPokemon()
	if err != nil {
		return err
	}
	
	helpers.PrintEncounter(pkmn)

	for {
		actionSelect := promptui.Select{
			Label: "What will you do?",
			Items: encounterActions,
		}
		_, action, err := actionSelect.Run()
		if err != nil {
			return err
		}

		err = handleSelectAction(pkmn, action)
		if err == ErrEmptyPrty {
			helpers.PrintEmptyParty()
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func handleSelectAction(pkmn *models.PokemonModel, action string) error {
	switch action {
	case "Catch":
		if err := handleCatch(pkmn); err != nil {
			return err
		}
		return nil
	case "Fight":
		if err := handleFight(pkmn); err != nil {
			return err
		}
		return nil
	case "Run":
		handleRun()
		return nil
	}
	return nil
}

func handleCatch(pkmn *models.PokemonModel) error {
	helpers.PrintCaught(pkmn)
	nicknameSelect := promptui.Select{
		Label: "Would you like to nickname your Pokemon?",
		Items: yesNo,
	}
	_, res, err := nicknameSelect.Run()
	if err != nil {
		return err
	}

	if res == "Yes" {
		nicknamePrompt := promptui.Prompt{
			Label: "Nickname",
			Default: helpers.CapitalizeName(pkmn.Name),
		}
		nickname, err := nicknamePrompt.Run()
		if err != nil {
			return err
		}
		pkmn.Nickname = nickname
	}
	if err = addToPartyOrBox(pkmn); err != nil {
		return err
	}
	return nil
}

func handleFight(foePkmn *models.PokemonModel) error {
	prtyPkmn := []models.PokemonModel{};
	if err := models.GetParty(&prtyPkmn); err != nil {
		return err
	}

	if len(prtyPkmn) == 0 {
		return ErrEmptyPrty
	}

	battlePkmn := prtyPkmn[rand.Intn(len(prtyPkmn))]
	if err := helpers.PrintBattle(&battlePkmn); err != nil {
		return err
	}
	helpers.PrintWildFoeFaint(foePkmn)
	// ready to evolve
	if battlePkmn.EvolvesTo == "" {
		return nil
	}
	helpers.PrintReadyToEvolve(&battlePkmn)
	// would you like to evolve
	evolveSelect := promptui.Select{
		Label: "Are you ready to evolve your Pokemon?",
		Items: yesNo,
	}
	_, res, err := evolveSelect.Run()
	if err != nil {
		return err
	}
	if res == "Yes" {
		evolvePokemon(&battlePkmn)
	}
	return nil;
}

func handleRun() {
	helpers.PrintRun()
}

func evolvePokemon(pkmn *models.PokemonModel) error {
	evolution, err := services.FetchPokemon(pkmn.EvolvesTo, pkmn.IsShiny)
	if err != nil {
		return err
	}
	preEvoPkmn := *pkmn
	if err = pkmn.EvolveInto(evolution); err != nil {
		return err
	}
	helpers.PrintEvolved(&preEvoPkmn, pkmn)
	return nil
}

func addToPartyOrBox(pkmn *models.PokemonModel) error {
	partyPkmn := []models.PokemonModel{}
	if err := models.GetParty(&partyPkmn); err != nil {
		return nil
	}

	if len(partyPkmn) < 6 {
		if err := addPokemon(pkmn, true); err != nil {
			return err
		}
		return nil
	}

	partySelect := promptui.Select{
		Label: "Would you like to add your new Pokemon to your party?",
		Items: yesNo,
	}
	_, res, err := partySelect.Run()
	if err != nil {
		return err
	}
	if res == "No" {
		if err := addPokemon(pkmn, false); err != nil {
			return err
		}
		return nil
	}

	partyNames := helpers.MapToString(partyPkmn, func(pkmn models.PokemonModel, index int) string {
		return helpers.FormattedPokemonName(&pkmn)
	})
	replaceSelect := promptui.Select{
		Label: "Your party is full! Choose a Pokemon to Replace:",
		Items: append(partyNames, "❌ Cancel"),
	}
	index, _, err := replaceSelect.Run()
	if err != nil {
		return err
	}

	if index == 6 {
		if err := addPokemon(pkmn, false); err != nil {
			return err
		}
		return nil
	}

	toReplace := &partyPkmn[index]
	if err = removePokemonFromParty(toReplace); err != nil {
		return err
	}

	if err = addPokemon(pkmn, true); err != nil {
		return err
	}
	return nil
}

func addPokemon(pkmn *models.PokemonModel, inParty bool) error {
	pkmn.IsInParty = inParty;
	if err := models.CreatePokemon(pkmn); err != nil {
		return err
	}
	if pkmn.IsInParty {
		helpers.PrintAddedToParty(pkmn)
		return nil
	}
	helpers.PrintAddedToBox(pkmn)
	return nil
}

func removePokemonFromParty(pkmn *models.PokemonModel) error {
	pkmn.IsInParty = false
	if err := models.UpdatePokemon(pkmn); err != nil {
		return err
	}
	helpers.PrintRemoved(pkmn)
	return nil
}