package main

import (
	"fmt"
	"log"
	"os"

	"github.com/7minutech/gator_go/internal/config"
)

func main() {
	myConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	myState := state{Config: &myConfig}

	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	myArgs := os.Args
	if len(myArgs) < 2 {
		fmt.Fprintln(os.Stderr, "error: not enough arguments were provided")
		os.Exit(1)
	}

	cmd := command{name: myArgs[1], args: myArgs[2:]}
	if err := cmds.run(&myState, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
