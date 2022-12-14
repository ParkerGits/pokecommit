package helpers

import (
	"fmt"

	"github.com/ParkerGits/pokecommit/models"
	"github.com/ParkerGits/pokecommit/services"
)

func PrintEncounter(pkmn models.PokemonModel) error {
	sprite, err := services.FetchAsciiSprite(pkmn.AsciiFrontSpriteUrl)
	if err != nil {
		return err
	}
	fmt.Println(sprite)
	fmt.Println(encounterText(pkmn) + exclamationMark + "\n")
	return nil
}

func PrintBattle(pkmn models.PokemonModel) error {
	if pkmn.AsciiBackSpriteUrl == "" {
		fmt.Println(FormattedPokemonName(pkmn) + space + usedText(pkmn) + exclamationMark);
		return nil;
	}

	sprite, err := services.FetchAsciiSprite(pkmn.AsciiBackSpriteUrl)
	if err != nil {
		return err
	}
	fmt.Println(sprite)
	fmt.Println(FormattedPokemonName(pkmn) + space + usedText(pkmn) + exclamationMark)
	return nil
}

func PrintWildFoeFaint(foePkmn models.PokemonModel) {
	fmt.Println(wildText(foePkmn) + space + faintText + period)
}

func PrintReadyToEvolve(pkmn models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + readyToEvolveText + exclamationMark)
}

func PrintCaught(pkmn models.PokemonModel) {
	fmt.Println(caughtText(pkmn) + exclamationMark + "\n")
}

func PrintStored(pkmn models.PokemonModel) {
	if pkmn.Nickname != "" {
		fmt.Println(formattedNickname(pkmn) + space + storedText)
		return
	}
	fmt.Println(formattedName(pkmn) + space + storedText)
}

func PrintView(pkmn models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn))
}

func PrintDeposited(pkmn models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + removedPartyText)
}

func PrintAddedToBox(pkmn models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + addedBoxText)
}

func PrintAddedToParty(pkmn models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + addedPartyText)
}

func PrintCatchOrFlee() {
	fmt.Println(catchOrFleeText)
}

func PrintEvolved(preEvo models.PokemonModel, evolution models.PokemonModel) {
	fmt.Println(yourPkmnText(preEvo) + space + evolvedText + space + formattedName(evolution) + exclamationMark)
}

func PrintRun() {
	fmt.Println(runText)
}

func PrintEmptyParty() {
	fmt.Println(emptyPartyText)
}

func PrintBadFilter() {
	fmt.Println(badFilterText)
}

func PrintNoPokemon() {
	fmt.Println(noPokemonText)
}

func PrintReplacedPokemon(replaced models.PokemonModel, new models.PokemonModel) {
	fmt.Println(FormattedPokemonName(replaced) + space + replacedText(new) + space + inYourPartyText + period)
}

func PrintCannotDepositLastPokemon() {
	fmt.Println(lastPokemonInPartyText)
}

func PrintNoPokemonInBox() {
	fmt.Println(noPokemonInBoxText)
}