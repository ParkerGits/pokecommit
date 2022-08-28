/*
Copyright © 2022 Parker Landon parkerjlandon@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ParkerGits/pokecommit/helpers"
	"github.com/ldez/go-git-cmd-wrapper/v2/commit"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
		message string
		rootCmd = &cobra.Command{
		Use:   "pokecommit",
		Short: "Catch a Pokemon on each commit!",
		Long: `
		PokeCommit is a CLI wrapper for git that lets you build a Pokemon team with each commit.
		Treat "pokecommit" as an alias for git commit. 
		
		pokecommit -m "Initial Commit"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(message) == 0 {
				return errors.New("Please enter a commit message")
			}
			output, err := git.Commit(commit.Message(message))
			if err != nil {
				return err
			}
			fmt.Println(output)
			if err = helpers.EngageRandomEncounter(); err != nil {
				return err
			}
			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pokegit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&message, "message", "m", "", "Add a message to your commit")
}


