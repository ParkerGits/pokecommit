package helpers

import (
	"fmt"
	"strings"

	"github.com/ParkerGits/pokecommit/models"
	"github.com/ParkerGits/pokecommit/services"
	"github.com/ttacon/chalk"
)

var (
	baseTextStyle = chalk.White.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.Black)
	encounterText = func(pkmn *models.PokemonModel) string {
		return fmt.Sprintf(baseTextStyle.Style("You encounter a %s"), GetTypeStyle(pkmn.Type1).Style(capitalizeName(pkmn.Name)))
	}
	exclamationMark = baseTextStyle.Style("!")
	period = baseTextStyle.Style(".")
	space = baseTextStyle.Style(" ")
	openParen = baseTextStyle.Style("(")
	shiny = baseTextStyle.Style("✨")
	closeParen = baseTextStyle.Style(")")
	caughtText = func(pkmn *models.PokemonModel) string {
		return fmt.Sprintf(baseTextStyle.Style("You caught the %s"), GetTypeStyle(pkmn.Type1).Style(capitalizeName(pkmn.Name)))
	}
	formattedName = func(pkmn *models.PokemonModel) string {
		return GetTypeStyle(pkmn.Type1).WithTextStyle(chalk.Bold).Style(capitalizeName(pkmn.Name))
	}
	formattedNickname = func(pkmn *models.PokemonModel) string {
		return GetTypeStyle(pkmn.Type1).WithTextStyle(chalk.Bold).Style(pkmn.Nickname)
	}
	storedText = baseTextStyle.Style("has been stored in your PC.")
	removedPartyText = baseTextStyle.Style("has been removed from your party.")
	addedPartyText = baseTextStyle.Style("has been added to your party!")
	addedBoxText = baseTextStyle.Style("has been added to your box.")
)

func PrintEncounter(pkmn *models.PokemonModel) error {
	sprite, err := services.FetchAsciiSprite(pkmn.AsciiSpriteUrl)
	if err != nil {
		return err
	}
	fmt.Println(sprite)
	fmt.Println(encounterText(pkmn) + exclamationMark + "\n")
	return nil
}

func PrintCaught(pkmn *models.PokemonModel) {
	fmt.Println(caughtText(pkmn) + exclamationMark + "\n")
}

func PrintStored(pkmn *models.PokemonModel) {
	if pkmn.Nickname != "" {
		fmt.Println(formattedNickname(pkmn) + space + storedText)
		return
	}
	fmt.Println(formattedName(pkmn) + space + storedText)
}

func PrintView(pkmn *models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn))
}

func PrintRemoved(pkmn *models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + removedPartyText)
}

func PrintAddedToBox(pkmn *models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + addedBoxText)
}

func PrintAddedToParty(pkmn *models.PokemonModel) {
	fmt.Println(FormattedPokemonName(pkmn) + space + addedPartyText)
}

func FormattedPokemonName(pkmn *models.PokemonModel) string {
	var fmtName strings.Builder
	if pkmn.Nickname != "" {
		fmtName.WriteString(formattedNickname(pkmn) + space + openParen + formattedName(pkmn) + closeParen)
	} else {
		fmtName.WriteString(formattedName(pkmn))
	}
	
	if pkmn.IsShiny {
		fmtName.WriteString(space + shiny)
	}
	return fmtName.String()
}

func capitalizeName(name string) string {
	return strings.Title(strings.ToLower(name))
}

func GetTypeStyle(pkType string) chalk.Style {
	switch pkType {
	case "normal":
		return chalk.White.NewStyle()
	case "fighting", "fire":
		return chalk.Red.NewStyle()
	case "ice", "flying":
		return chalk.Cyan.NewStyle()
	case "psychic", "fairy", "poison":
		return chalk.Magenta.NewStyle()
	case "water", "dragon":
		return chalk.Blue.NewStyle()
	case "ghost", "steel", "dark":
		return chalk.Dim.NewStyle()
	case "electric", "bug", "ground", "rock":
		return chalk.Yellow.NewStyle()
	case "grass":
		return chalk.Green.NewStyle()
	}
	return nil
}