package main

import "fmt"

func commandMapB(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)

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