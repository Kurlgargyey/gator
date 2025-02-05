package main

import (
	"context"
	"net/http"
)


type RSSFeed struct {
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	res, err := http.NewRequestWithContext(ctx, "get", feedURL)

	return nil, nil
}