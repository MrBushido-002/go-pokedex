package main

import (
	"fmt"
    "github.com/MrBushido-002/go-pokedex/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	res := pokeapi.LocationAreaResponse{}
	var err error
	if cfg.Next == nil {
		res, err = cfg.pokeapiClient.GetAreaLocation(nil)
		if err != nil {
			return err
		}
	} else {
		res, err = cfg.pokeapiClient.GetAreaLocation(cfg.Next)
		if err != nil {
			return err
		}
	}
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for i := 0;i < len(res.Results); i++ {
		fmt.Println(res.Results[i].Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	res := pokeapi.LocationAreaResponse{}
	var err error
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
			return nil
		}

	res, err = cfg.pokeapiClient.GetAreaLocation(cfg.Previous)
	if err != nil {
		return err
	}
	
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	
	for i := 0;i < len(res.Results); i++ {
		fmt.Println(res.Results[i].Name)
	}
	return nil
}