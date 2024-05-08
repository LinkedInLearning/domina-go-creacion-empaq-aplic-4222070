package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Metadata struct {
	Author  string
	Version string
	Date    string
	Commit  string
}

var metadata Metadata

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pkm",
	Short: "A CLI tool for managing Pokemon data",
	Long:  `A CLI tool for managing Pokemon data`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf(`Welcome to the Pokémon CLI %s, created with 💖 by %s!
Commit Hash: %s
Build Date: %s`, metadata.Version, metadata.Author, metadata.Commit, metadata.Date)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(m Metadata) {
	metadata = m
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
