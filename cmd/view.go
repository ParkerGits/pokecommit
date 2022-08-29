/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
	"github.com/spf13/cobra"
)

var filter string

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
		if len(filter) == 0 {
			for _, pkmn := range allPkmn {
				helpers.PrintView(&pkmn)
			}
			return nil
		}
		for _, pkmn := range allPkmn {
				if(strings.Contains(pkmn.Name, filter) || strings.Contains(pkmn.Nickname, filter)) {
					helpers.PrintView(&pkmn)
				}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter your Pokemon by name or nickname.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
