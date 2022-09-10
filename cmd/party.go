/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ParkerGits/pokecommit/handlers"
	"github.com/ParkerGits/pokecommit/models"
	"github.com/spf13/cobra"
)

var add bool
var deposit bool

// partyCmd represents the party command
var partyCmd = &cobra.Command{
	Use:   "party",
	Short: "Manage your party.",
	Long: `View and edit your party!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		allPkmn := []models.PokemonModel{}
		if err := models.GetAllPokemon(&allPkmn); err != nil {
				return err
		}

		if deposit {
			return handlers.HandleDepositPokemon(&allPkmn)
		}

		if add {
			return handlers.HandleAddPokemon(&allPkmn)
		}
		
		handlers.PrintPartyPokemon(&allPkmn)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(partyCmd)
	partyCmd.Flags().BoolVarP(&add, "add", "a", false, "Add a Pokemon in your box to your party")
	partyCmd.Flags().BoolVarP(&deposit, "deposit", "d", false, "Deposit a Pokemon from your party into your box")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// partyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// partyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
