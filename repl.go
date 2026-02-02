package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrBushido-002/go-pokedex/internal/pokeapi"
)


type config struct {
	pokeapiClient pokeapi.Client
	Next *string
	Previous *string
	caughtPokemon map[string]pokeapi.PokemonInfo
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("Pokedex >")
		scan := scanner.Scan()
		if scan == true {
			txt := scanner.Text()
			clean_txt := cleanInput(txt)
			if len(clean_txt) == 0 {
				continue
			}
			cmd := clean_txt[0]
			args := clean_txt[1:]

			command, ok := getCommands()[cmd]
			
			if ok {
				err := command.callback(cfg, args)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
		
	}
	return
}

func cleanInput(text string) ([]string) {
	lower := strings.ToLower(text)
	trim := strings.Trim(lower, " ")
	split := strings.Fields(trim)
	return split
}

type cliCommand struct {
	Name string
	description string
	callback func(*config, []string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			Name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			Name:        "explore",
			description: "Explore a location area for Pokémon",
			callback:    commandExplore,
		},
		"catch": {
			Name: "catch",
			description: "Throw a pokeball for a chance to catch it and add it to your pokedex",
			callback: commandCatch,
		},
		"inspect": {
			Name: "inspect",
			description: "inspect a pokemon if it has already been caught",
			callback: commandInspect,
		},
		"pokedex": {
			Name: "pokedex",
			description: "lists all pokemon caught",
			callback: commandPokedex,
		},

	}
}
