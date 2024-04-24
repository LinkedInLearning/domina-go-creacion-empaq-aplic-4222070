package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(name string) (Pokemon, error) {
	var pokemon Pokemon // zero value of Pokemon struct
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		return pokemon, fmt.Errorf("failed to get pokemon: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return pokemon, nil
}

func GetTypes() (TypesResult, error) {
	var results TypesResult // zero value of TypesResult struct
	resp, err := http.Get("https://pokeapi.co/api/v2/type")
	if err != nil {
		return results, fmt.Errorf("failed to get types: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return results, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return results, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return results, nil
}
