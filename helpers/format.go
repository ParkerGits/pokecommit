package helpers

import (
	"fmt"
	"strings"

	"github.com/ParkerGits/pokecommit/models"
	"github.com/ttacon/chalk"
)

func PrintPkmn(pkmn models.PokemonModel, sprite string) {
	fmt.Println(sprite)
	encounterTxt := GetTypeStyle(pkmn.Type1).Style(capitalizeName(pkmn.Name))
	fmt.Printf(chalk.White.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.Black).Style("You encountered a %s\n"), encounterTxt)
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