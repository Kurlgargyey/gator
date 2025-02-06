package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		usr, err := s.db.GetUser(context.Background(), s.cfg.CurrUser)
		if err != nil {
			return fmt.Errorf("error retrieving current user: %w", err)
		}
		return handler(s, c, usr)
	}
}