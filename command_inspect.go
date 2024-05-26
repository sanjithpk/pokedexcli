package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return fmt.Errorf("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}
