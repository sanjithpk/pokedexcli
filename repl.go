package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapB,
		},
	}
}


func repl(cfg *config)  {
	for {
		var command string
		fmt.Print("Pokedex > ")
		fmt.Scanf("%s", &command)
		commands := getCommands()
		if _, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := commands[command].callback(cfg)
		if err != nil{
			fmt.Println(err)
			continue
		}
	}
}