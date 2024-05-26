package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists pokemons in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch pokemons in the area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect your pokemons",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func getWords(words string) []string {
	wordList := strings.Fields(strings.ToLower(words))
	return wordList
}

func repl(cfg *config) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pokedex > ")
		input, _ := reader.ReadString('\n')
		commands := getCommands()
		words := getWords(input)
		command := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		if _, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := commands[command].callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
