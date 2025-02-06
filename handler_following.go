package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: following")
	}

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrUser)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), usr.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}
	fmt.Printf("Follows for user %s:\n", usr.Name)
	for _, follow := range follows {
		fmt.Printf("Feed: %s\n", follow.FeedName)
	}

	return nil
}