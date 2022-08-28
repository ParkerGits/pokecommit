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
			Label: "Would you like to nickname your pokemon?",
			Items: yesNo,
		}
		_, res, err := nicknameSelect.Run()
		if err != nil {
			return err
		}

		if res == "No" {
			if err := models.CreatePokemon(pkmn); err != nil {
				return err
			}
			PrintStored(pkmn)
			return nil
		}
		
		nicknamePrompt := promptui.Prompt{
			Label: "Nickname",
			Default: capitalizeName(pkmn.Name),
		}
		nickname, err := nicknamePrompt.Run()
		if err != nil {
			return err
		}
		pkmn.Nickname = nickname
		PrintStored(pkmn)
		if err = models.CreatePokemon(pkmn); err != nil {
			return err
		}
		return nil
	}
	return nil
}

