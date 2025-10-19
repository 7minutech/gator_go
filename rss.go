package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/7minutech/gator_go/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	var feed RSSFeed

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &feed, fmt.Errorf("error: creating request for feed url: %w", err)
	}

	req.Header.Set("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &feed, fmt.Errorf("error: sending feed request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &feed, fmt.Errorf("error: bad status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &feed, fmt.Errorf("error: reading response body of feed status code (%d): %w", resp.StatusCode, err)
	}

	if err := xml.Unmarshal(data, &feed); err != nil {
		return &feed, fmt.Errorf("error: unmarshaling response data of feed: %w", err)
	}

	return &feed, nil
}

func decodeEscapedHtml(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}

func parsePublishTime(s string) time.Time {
	layouts := []string{
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t.UTC()
		}
	}
	return time.Now().UTC()
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error: getting next feed to fetch: %w", err)
	}

	if err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("error: marking next feed as fetched: %w", err)
	}
	fmt.Printf("Fetching: %s (%s)\n", nextFeed.Name, nextFeed.Url)

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error: fetching RSS feed: %w", err)
	}

	decodeEscapedHtml(rssFeed)

	ctx := context.Background()

	for _, item := range rssFeed.Channel.Item {
		now := time.Now().UTC()

		title := sql.NullString{String: item.Title, Valid: item.Title != ""}
		desc := sql.NullString{String: item.Description, Valid: item.Description != ""}

		publishTime := parsePublishTime(item.PubDate)
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: publishTime,
			FeedID:      nextFeed.ID,
		}

		if _, err := s.db.CreatePost(ctx, postParams); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			log.Printf("error: createing post: %v\n", err)
		}
	}

	return nil
}
