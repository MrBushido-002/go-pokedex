package main

import(
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	fmt.Printf(
	`Welcome to the Pokedex!
	Usage:
	
	help: Displays a help message
	exit: Exit the Pokedex
	map: Get the next page of 20 locations
	mapb: Get the previous 20 locations
	explore: Explore a location and Display pokemon found there
	catch: Throw a pokeball for a chance to catch it and add it to your pokedex
	inspect: Inspect a pokemon if it has been caught
	pokedex: Lists all caught pokemon
	`)
	return nil
}
