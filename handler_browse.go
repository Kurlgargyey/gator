package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: browse [<limit>] ")
	}
	var PostLimit int

	if len(cmd.args) == 0 {
		PostLimit = 2
	} else {
		l, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Println("you did not provide a valid limit... setting limit to 2.")
			PostLimit = 2
		} else {
			PostLimit = l
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:    user.ID,
		PostLimit: int32(PostLimit),
	})

	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		feed, _ := s.db.GetFeedByID(context.Background(), post.FeedID)
		fmt.Printf("Post: %s\n Feed: %s\n Published: %s\n", post.Title, feed.Name, post.PublishedAt.Format("Mon, 02.01.2006 15:04"))
	}
	return nil
}
