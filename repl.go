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
	Next string
	Previous string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)


	for i := 0; ;i++ {
		fmt.Print("Pokedex >")
		scan := scanner.Scan()
		if scan == true {
			txt := scanner.Text()
			clean_txt := cleanInput(txt)
			if len(clean_txt) == 0 {
				continue
			}
			cmd := clean_txt[0]

			command, ok := getCommands()[cmd]
			
			if ok {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Printf("Unknown command")
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
	callback func(*config) error
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
	}
}
