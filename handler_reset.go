package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("usage: reset")
	}
	s.db.DeleteUsers(context.Background())
	s.cfg.SetUser("")
	fmt.Println("Database was reset.")
	return nil
}
