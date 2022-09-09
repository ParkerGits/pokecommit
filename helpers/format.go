package helpers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ParkerGits/pokecommit/models"
	"github.com/ttacon/chalk"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	titleCaser = cases.Title(language.English)
	baseTextStyle = chalk.White.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.Black)
	encounterText = func(pkmn *models.PokemonModel) string {
		return fmt.Sprintf(baseTextStyle.Style("You encounter a %s"), GetTypeStyle(pkmn.Type1).Style(CapitalizeName(pkmn.Name)))
	}
	exclamationMark = baseTextStyle.Style("!")
	period = baseTextStyle.Style(".")
	space = baseTextStyle.Style(" ")
	openParen = baseTextStyle.Style("(")
	shiny = baseTextStyle.Style("✨")
	party = baseTextStyle.Style("⭐️")
	closeParen = baseTextStyle.Style(")")
	caughtText = func(pkmn *models.PokemonModel) string {
		return fmt.Sprintf(baseTextStyle.Style("You caught the %s"), GetTypeStyle(pkmn.Type1).Style(CapitalizeName(pkmn.Name)))
	}
	formattedName = func(pkmn *models.PokemonModel) string {
		return GetTypeStyle(pkmn.Type1).WithTextStyle(chalk.Bold).Style(CapitalizeName(pkmn.Name))
	}
	formattedNickname = func(pkmn *models.PokemonModel) string {
		return GetTypeStyle(pkmn.Type1).WithTextStyle(chalk.Bold).Style(pkmn.Nickname)
	}
	usedText = func(pkmn *models.PokemonModel) string {
		randMove := GetRandPokemonMove(pkmn)
		return fmt.Sprintf(baseTextStyle.Style("used %s"), GetTypeStyle(pkmn.Type1).WithTextStyle(chalk.Bold).Style(titleCaser.String(randMove)))
	}
	wildText = func(pkmn *models.PokemonModel) string { 
		return fmt.Sprintf(baseTextStyle.Style("The wild %s"), FormattedPokemonName(pkmn))
	}
	yourPkmnText = func(pkmn *models.PokemonModel) string {
		return fmt.Sprintf(baseTextStyle.Style("Your %s"), FormattedPokemonName(pkmn))
	}
	faintText = baseTextStyle.Style("has fainted")
	readyToEvolveText = baseTextStyle.Style("is ready to evolve")
	storedText = baseTextStyle.Style("has been stored in your PC.")
	removedPartyText = baseTextStyle.Style("has been removed from your party.")
	addedPartyText = baseTextStyle.Style("has been added to your party!")
	addedBoxText = baseTextStyle.Style("has been added to your box.")
	partyEmptyText = baseTextStyle.Style("Your party is empty! You must capture this Pokemon or flee.")
	evolvedText = baseTextStyle.Style("has evolved into")
	runText = baseTextStyle.Style("You flee the battle...")
	pokemonMoves = map[string][3]string{
		"bug": {"bug buzz", "silver wind", "x-scissor"},
		"dark": {"dark pulse", "knock off", "night slash"},
		"dragon": {"dragon claw", "dragon pulse", "draco meteor"},
		"electric": {"spark", "discharge", "thunder"},
		"fighting": {"brick break", "focus blast", "close combat"},
		"fire": {"flamethrower", "flare blitz", "fire punch"},
		"flying": {"brave bird", "aerial ace", "hurricane"},
		"ghost": {"shadow ball", "shadow sneak", "shadow claw"},
		"grass": {"giga drain", "seed bomb", "power whip"},
		"ground": {"earth power", "mud bomb", "earthquake"},
		"ice": {"ice beam", "ice shard", "blizzard"},
		"normal": {"return", "quick attack", "hyper beam"},
		"poison": {"gunk shot", "sludge bomb", "poison jab"},
		"psychic": {"psychic", "zen headbutt", "confusion"},
		"rock": {"stone edge", "ancientpower", "rock tomb"},
		"steel": {"iron head", "flash cannon", "iron tail"},
		"water": {"water pulse", "hydro pump", "waterfall"},
	}
)

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

	if pkmn.IsInParty {
		fmtName.WriteString(space + party)
	}
	return fmtName.String()
}

func CapitalizeName(name string) string {
	return titleCaser.String(name)
}

func GetRandPokemonMove(pkmn *models.PokemonModel) string {
	randMoveIdx := rand.Intn(3)
	return pokemonMoves[pkmn.Type1][randMoveIdx]
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