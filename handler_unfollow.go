package main

import (
	"fmt"
	"gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("")
	}

	feedUrl := cmd.args[0]
}