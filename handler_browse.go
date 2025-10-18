package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/7minutech/gator_go/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("error: browse expects only 1 optional arg")
	}

	limit := "2"

	if len(cmd.args) == 1 {
		limit = cmd.args[0]
	}

	parsedLimit, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return fmt.Errorf("error: parsing limit: %w", err)
	}

	postsForUserParams := database.GetPostsForUserParams{UserID: user.ID, Limit: int32(parsedLimit)}

	posts, err := s.db.GetPostsForUser(context.Background(), postsForUserParams)
	if err != nil {
		return fmt.Errorf("error: getting posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s (%s)\n\tDesc: %s\n\n", post.Title.String, post.Description.String, post.Url)
	}

	return nil
}
