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
	fmt.Println("fetching feed...")
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching next feed: %w", err)
	}
	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	fmt.Printf("fetching posts from feed %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		pubDate, err := parsePubDate(item.PubDate)
		if err != nil {
			return fmt.Errorf("bad pubdate format: %s. error: %w", item.PubDate, err)
		}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: parseDescription(item.Description),
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
			continue
		}
		if err != nil {
			return fmt.Errorf("error creating post: %w", err)
		}
		fmt.Printf("created post: %s\n", post.Title)
	}

	return nil
}

func parsePubDate(dateString string) (time.Time, error) {
	_ = "Mon, 02 Jan 2006 15:04:05 +0000"
	return time.Parse(time.RFC1123Z, dateString)
}

func parseDescription(description string) sql.NullString {
	var sqlDescription sql.NullString
	if len(strings.TrimSpace(description)) == 0 {
		sqlDescription = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		sqlDescription = sql.NullString{
			String: strings.TrimSpace(description),
			Valid:  true,
		}
	}
	return sqlDescription
}
