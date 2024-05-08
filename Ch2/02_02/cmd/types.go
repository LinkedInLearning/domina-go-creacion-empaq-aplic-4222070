package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AlecAivazis/survey/v2"

	"github.com/linkedinlearning/domina-go/binaries/pokemon"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get details about a Pokémon",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := getTypes()
		if err != nil {
			log.Fatalf("failed to get types: %v", err)
		}

		types := make([]string, len(ts))
		for i, t := range ts {
			types[i] = t.Name
		}

		t := ""
		prompt := &survey.Select{
			Message: "¿Qué tipo de Pokémon buscas?",
			Options: types,
		}
		survey.AskOne(prompt, &t)

		pokemons, err := getPokemonsByType(t)
		if err != nil {
			log.Fatalf("failed to get pokemon: %v", err)
		}

		p := ""
		prompt = &survey.Select{
			Message: "Elige un Pokémon de la lista:",
			Options: pokemons,
		}
		survey.AskOne(prompt, &p)

		if err := getPokemon(p); err != nil {
			log.Fatalf("failed to get pokemon: %v", err)
		}
	},
}

func getTypes() ([]pokemon.Type, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/type")
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var search pokemon.TypeSearch
	json.Unmarshal(body, &search)

	return search.Results, nil
}

func getPokemonsByType(t string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/type/%s", t))
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var search pokemon.PokemonByType
	json.Unmarshal(body, &search)

	var pokemons []string
	for _, p := range search.Pokemon {
		pokemons = append(pokemons, p.Pokemon.Name)
	}

	return pokemons, nil
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
	fmt.Printf("%s, I choose you!\n", pokemon.Name)

	return nil
}
