package main

import "fmt"

func handlerHelp(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: help expects 0 args")
	}

	for cmd, desc := range cmdDescriptions {
		fmt.Printf("Command: %s", cmd)
		fmt.Println(desc)
	}

	return nil
}
