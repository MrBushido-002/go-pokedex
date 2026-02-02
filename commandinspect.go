package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args[] string) error {
	if len(args) != 1 {
		return errors.New("no pokemon given")
	}

	pokemonName := args[0]

	res, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemon not caught yet")
	}
	fmt.Printf("Name: %s\n", res.Name)
	fmt.Printf("Height: %v\n", res.Height)
	fmt.Printf("Weight: %v\n", res.Weight)
	fmt.Printf("Stats:\n")

	for _, val := range res.Stats {
		fmt.Printf("	- %v: %v\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types")
	for _, val := range res.Types {
		fmt.Printf("	- %v\n", val.Type.Name)
	}
	return nil
}