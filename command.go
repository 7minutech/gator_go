package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s.cfg == nil {
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
