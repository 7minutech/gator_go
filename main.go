package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/7minutech/gator_go/internal/config"
	"github.com/7minutech/gator_go/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	myConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", myConfig.DbUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening db: %w", err)
		os.Exit(1)
	}
	dbQuries := database.New(db)

	myState := state{cfg: &myConfig, db: dbQuries}

	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("help", handlerHelp)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("error: no current user set")
		}

		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error: fetching user for middleware: %w", err)
		}

		return handler(s, cmd, currentUser)
	}
}
