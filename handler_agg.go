package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("error: agg expects a single argument, the url")
	}

	testUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		return fmt.Errorf("error: fetching feed: %w", err)
	}

	decodeEscapedHtml(feed)

	fmt.Printf("%+v\n", feed)
	return nil
}
