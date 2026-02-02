package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:\n")
	for _, val := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", val.Name)
	}
	return nil
}