package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemons in Pokedex")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemonName)
	}

	return nil
}
