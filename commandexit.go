package main

import(
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Printf(
	`Welcome to the Pokedex!
	Usage:
	
	help: Displays a help message
	exit: Exit the Pokedex
	`)
	return nil
}