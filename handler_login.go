package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("usage: login <username>")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		log.Fatal("user doesn't exist!")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("user has been set to %s.\n", cmd.args[0])
	return nil
}