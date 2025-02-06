package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}

	feedUrl := cmd.args[0]
	_, err := url.Parse(feedUrl)
	if err != nil {
		return fmt.Errorf("url argument did not contain a valid url: %s. encountered the following error: %w", feedUrl, err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating follow: %w", err)
	}
	fmt.Printf("New follow: %s\nUser: %s\n", follow.FeedName, follow.UserName)
	return nil
}