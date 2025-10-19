package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("error: agg expects a single argument, the time_between_reqs")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("scrape error: %v\n", err)
			continue
		}
	}

	return nil
}
