package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrUser)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	feedUrl := cmd.args[1]

	_, err = url.Parse(feedUrl)
	if err != nil {
		return fmt.Errorf("url argument did not contain a valid url: %s. encountered the following error: %w", feedUrl, err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: feedUrl,
		UserID: usr.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}