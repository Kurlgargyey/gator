package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"net/url"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: unfollow <feed-url>")
	}

	feedUrl := cmd.args[0]
	if _, err := url.Parse(feedUrl); err != nil {
		return fmt.Errorf("feed url was not a valid url: %s. error: %w", feedUrl, err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}
	s.db.RemoveFollow(context.Background(), database.RemoveFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	return nil
}