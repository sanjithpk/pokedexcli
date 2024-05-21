package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Pokemon in %s:\n", locationAreaName)
	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
