package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("no location given")
	}
	
	locationName := args[0]

	res, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", res.Name)
	fmt.Println("Found Pokemon:")

	for _, enc := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}