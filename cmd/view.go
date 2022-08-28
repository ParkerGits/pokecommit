/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"


	"github.com/ParkerGits/pokecommit/models"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your captured Pokemon",
	Long: `View all the Pokemon captured through your commits.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		allPkmn := []models.PokemonModel{}
		if err := models.GetAllPokemon(&allPkmn); err != nil {
				return err
		}
		for _, pkmn := range allPkmn {
			fmt.Println(pkmn.Name)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
