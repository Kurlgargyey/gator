package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	fmt.Println("Logging in...")
	if len(cmd.args) == 0 {
		return errors.New("usage: login <username>")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("user has been set to %s.\n", cmd.args[0])
	return nil
}