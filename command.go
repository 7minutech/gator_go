package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/7minutech/gator_go/internal/database"
	"github.com/google/uuid"
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return fmt.Errorf("error: login expects a single argument, the username")
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("error: %s does not exist", name)
		}
		return fmt.Errorf("error: fetching %s from users: %w", name, err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("error: setting login for user %s: %w", name, err)
	}

	fmt.Printf("user %s has been set\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return fmt.Errorf("error: register expects a single argument, the name")
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("error: %s is already registered", name)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error: checking if name: %s exists in users: %w", name, err)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}
	newUser, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error: creating user %s: %w", name, err)
	}

	if err := s.cfg.SetUser(newUser.Name); err != nil {
		return fmt.Errorf("error: registering user %s: %w", name, err)
	}

	fmt.Printf("user %s has been created\n", newUser.Name)
	fmt.Printf("%+v\n", newUser)

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: reset expects zero arguments")
	}

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("error: deleting users: %w", err)
	}

	fmt.Println("users were successfully reset")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: users expects zero arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: fetching users: %w", err)
	}

	for _, user := range users {
		msg := user.Name
		if user.Name == s.cfg.CurrentUserName {
			msg += " (current)"
		}
		fmt.Println(msg)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	testUrl := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		return fmt.Errorf("error: fetching feed: %w", err)
	}

	decodeEscapedHtml(feed)

	fmt.Printf("%+v\n", feed)
	return nil
}
