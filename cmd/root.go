/*
Copyright © 2022 Parker Landon parkerjlandon@gmail.com
*/
package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
		toggle string
		rootCmd = &cobra.Command{
		Use:   "pokecommit",
		Short: "Catch a Pokemon on each commit!",
		Long: `PokeCommit is a CLI wrapper for Git that lets you build a Pokemon team with each commit.`,
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
	rand.Seed(time.Now().UnixNano())
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pokegit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}


