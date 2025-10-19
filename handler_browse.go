package main

import (
	"context"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"strings"

	"github.com/7minutech/gator_go/internal/database"
)

var (
	brRe   = regexp.MustCompile(`(?i)<\s*br\s*/?>`)
	pOpen  = regexp.MustCompile(`(?i)<\s*p[^>]*>`)
	pClose = regexp.MustCompile(`(?i)</\s*p\s*>`)
	tagRe  = regexp.MustCompile(`<[^>]+>`)
)

func plainText(html string) string {
	s := brRe.ReplaceAllString(html, "\n")
	s = pOpen.ReplaceAllString(s, "\n")
	s = pClose.ReplaceAllString(s, "\n")
	s = tagRe.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	return s
}

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
		title := (post.Title.String)
		url := post.Url
		desc := html.UnescapeString(plainText(post.Description.String))
		fmt.Printf("%s (%s)\n\tDesc: %s\n\n", title, url, desc)
	}

	return nil
}
