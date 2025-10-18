package main

import (
	"context"
	"fmt"
	"time"

	"github.com/7minutech/gator_go/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("error: addfeed expects exactly 2 arguments the name and url")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error: getting current user for feed creation: %w", err)
	}

	name := cmd.args[0]
	url := cmd.args[1]
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error: creating %s feed: %w", name, err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	if _, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams); err != nil {
		return fmt.Errorf("error: creating feed follow while creating feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: feeds expects zero arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error: getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s | Feed url: %s | Created by: %s\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}
