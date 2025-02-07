package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("  - URL: %s\n", feed.Url)
		fmt.Printf("  - Owner: %s\n", user.Name)
	}
	return nil
}
