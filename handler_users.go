package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("usage: users")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		line := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrUser {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}