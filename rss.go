package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"html"
	"io"
	"net/http"
	"strings"
	"time"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.DefaultClient

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(strings.ToLower(contentType), "xml") {
		return nil, fmt.Errorf("unexpected content type: %s", contentType)
	}

	xmlRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading XML: %w", err)
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(xmlRaw, &rssFeed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling XML: %w", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for _, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed URL: %w", err)
	}
	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching next feed: %w", err)
	}
	fmt.Printf("Items in feed %s:\n", feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		if i > 4 {
			break
		}
		fmt.Printf("  - %s\n", item.Title)
	}

	return nil
}
