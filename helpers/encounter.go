package helpers

import (
	"github.com/ParkerGits/pokecommit/services"
	"github.com/ParkerGits/pokecommit/models"
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
)

func EngageRandomEncounter() error {
	pkmn, err := services.FetchRandomPokemon()
	if err != nil {
		return err
	}
	
	PrintEncounter(&pkmn)

	actionSelect := promptui.Select{
		Label: "What will you do?",
		Items: encounterActions,
	}
	_, action, err := actionSelect.Run()
	if err != nil {
		return err
	}

	if err = handleSelectAction(&pkmn, action); err != nil {
		return err
	}
	
	return nil
}

func handleSelectAction(pkmn *models.PokemonModel, action string) error {
	switch action {
	case "Catch":
		PrintCaught(pkmn)
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
				Default: capitalizeName(pkmn.Name),
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

	partyNames := MapToString(partyPkmn, func(pkmn models.PokemonModel, index int) string {
		return FormattedPokemonName(&pkmn)
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
		PrintAddedToParty(pkmn)
		return nil
	}
	PrintAddedToBox(pkmn)
	return nil
}

func removePokemonFromParty(pkmn *models.PokemonModel) error {
	pkmn.IsInParty = false
	if err := models.UpdatePokemon(pkmn); err != nil {
		return err
	}
	PrintRemoved(pkmn)
	return nil
}