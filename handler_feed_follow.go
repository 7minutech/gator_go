package main

import (
	"context"
	"fmt"
	"time"

	"github.com/7minutech/gator_go/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: follow expects one argument the url")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error: fetching current user while following feed: %w", err)
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error: fetching feed while following feed")
	}

	feedFollowArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), feedFollowArgs)
	if err != nil {
		return fmt.Errorf("error: creeating feed follow: %w", err)
	}

	fmt.Printf("user: %s\nfeed: %s\n", feedFollowRow.UserName, feedFollowRow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: following expects zero arguments")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error: fetching current user while showing following: %w", err)
	}

	followings, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error: fetching follows for current user: %w", err)
	}

	for _, follow := range followings {
		fmt.Printf("Feed: %s\n", follow.FeedName)
	}

	return nil

}
