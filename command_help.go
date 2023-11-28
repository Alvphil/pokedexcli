package main

import "fmt"

func commandHelp() error {
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("%v: %v \n", command.name, command.description)
	}
	return nil
}
