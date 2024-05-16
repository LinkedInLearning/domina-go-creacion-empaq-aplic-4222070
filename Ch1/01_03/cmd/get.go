package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/linkedinlearning/domina-go/clis/pokemon"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get details about a Pokémon",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := getPokemon(args[0]); err != nil {
			log.Fatalf("failed to get pokemon: %v", err)
		}
	},
}

func getPokemon(name string) error {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		return fmt.Errorf("failed to get pokemon: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var pokemon pokemon.Pokemon
	json.Unmarshal(body, &pokemon)
	fmt.Printf("Name: %s, %+v\n", pokemon.Name, pokemon)

	return nil
}
