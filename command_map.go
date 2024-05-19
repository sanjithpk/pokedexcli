package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}