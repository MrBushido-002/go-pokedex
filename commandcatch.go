package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
)


func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("no pokemon given")
	}

	pokemonName := args[0]

	res, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	random_int := rand.IntN(1000)
	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)

	if random_int >= res.BaseExperience {
		fmt.Printf("%s was caught!\n", res.Name)
		cfg.caughtPokemon[res.Name] = res
		return nil
	}

	fmt.Printf("%s escaped!\n", res.Name)

	return nil

}