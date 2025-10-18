package main

import (
	"context"
	"fmt"
	"time"

	"github.com/7minutech/gator_go/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: follow expects one argument the url")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error: fetching feed while following feed")
	}

	now := time.Now().UTC()
	feedFollowArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
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

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: following expects zero arguments")
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

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: unfollow expects exactly one argument the url of the feed")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error: fetching feed for unfollow: %w", err)
	}

	deleteParams := database.DeleteFeedFollowParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}
	if err := s.db.DeleteFeedFollow(context.Background(), deleteParams); err != nil {
		return fmt.Errorf("error: deleting feed follow: %w", err)
	}

	return nil
}
