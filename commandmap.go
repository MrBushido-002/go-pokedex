package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	res := LocationAreaResponse{}
	var err error
	if cfg.Next == "" {
		res, err = GetAreaLocation("")
		if err != nil {
			return err
		}
	} else {
		res, err = GetAreaLocation(cfg.Next)
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

func commandMapb(cfg *config) error {
	res := LocationAreaResponse{}
	var err error
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
			return nil
		}

	res, err = GetAreaLocation(cfg.Previous)
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