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

var cmdDescriptions = map[string]string{
	"help":      "\n\tUsage: help\n\tDescription: shows list of commands and their descriptions",
	"login":     "\n\tUsage: login <username>\n\tDescription: switch the current user",
	"register":  "\n\tUsage: register <username>\n\tDescription: create a new user and login as that user",
	"reset":     "\n\tUsage: reset\n\tDescription: removes all records from all tables",
	"users":     "\n\tUsage: users\n\tDescription: displays all registered users and current user",
	"agg":       "\n\tUsage: agg <duration string>\n\tDescription: aggregate feeds continously and insert newly created posts\n\tmeant to run in a seperate shell",
	"addfeed":   "\n\tUsage: addfeed <name> <url>\n\tDescription: creates a new feed",
	"follow":    "\n\tUsage: follow <feed url>\n\tDescription: current user follows specified feed",
	"following": "\n\tUsage: following\n\tDescription: shows all feeds followed by current user",
	"unfollow":  "\n\tUsage: unfollow <feed url>\n\tDescription: current user unfollows specified feed",
	"browse":    "\n\tUsage: browse [limit]\n\tDescription: shows current user's followed feed posts with optional limit arg",
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
