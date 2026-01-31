package main
import (
	"strings"
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)


type config struct {
	pokeapiClient pokeapi.Client
	Next string
	Previous string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{}

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

			command, ok := cmdmap[cmd]
			
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
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
