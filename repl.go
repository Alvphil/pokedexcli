package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()
	for {
		fmt.Print("pokedex > ")
		if !scanner.Scan() {
			break
		}
		commandText := scanner.Text()
		commandCleaned := cleanInput(commandText)
		commandName := commandCleaned[0]
		command, ok := availableCommands[commandName]
		if ok {
			command.callback()
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
