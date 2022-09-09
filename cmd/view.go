/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ParkerGits/pokecommit/handlers"
	"github.com/ParkerGits/pokecommit/models"
	"github.com/spf13/cobra"
)

var filter string
var party bool

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your captured Pokemon",
	Long: `
	View all the Pokemon captured through your commits.
	User the --filter (-f) flag to filter your Pokemon
	by name or nickname.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		allPkmn := []models.PokemonModel{}
		if err := models.GetAllPokemon(&allPkmn); err != nil {
				return err
		}

		if party {
			handlers.PrintPartyPokemon(&allPkmn)
			return nil
		}

		if len(filter) == 0 {
			handlers.PrintAllPokemon(&allPkmn);
			return nil
		}

		handlers.PrintFilteredPokemon(&allPkmn, filter)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().BoolVarP(&party, "party", "p", false, "View your party Pokemon.")
	viewCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter your Pokemon by name or nickname.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
