/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ParkerGits/pokecommit/models"
	"github.com/spf13/cobra"
)

// partyCmd represents the party command
var partyCmd = &cobra.Command{
	Use:   "party",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		partyPkmn := []models.PokemonModel{}
		if err := models.GetParty(&partyPkmn); err != nil {
			return err
		}
		if len(partyPkmn) == 0 {
			return errors.New("Party is empty!")
		}
		for _, pkmn := range partyPkmn {
			helpers.PrintView(&pkmn)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(partyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// partyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// partyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
