package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: following")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}
	fmt.Printf("Follows for user %s:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("Feed: %s\n", follow.FeedName)
	}

	return nil
}