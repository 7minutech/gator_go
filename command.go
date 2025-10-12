package main

import (
	"fmt"

	"github.com/7minutech/gator_go/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s.Config == nil {
		return fmt.Errorf("error: state for command to run does not exists")
	}
	f, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("error: unknown command %s", cmd.name)
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error: login expects a single argument, the username")
	}
	if err := s.Config.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", s.Config.CurrentUserName)
	return nil
}
