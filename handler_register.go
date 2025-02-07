package main

import (
	"context"
	"errors"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("usage: register <username>")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		log.Fatal(errors.New("username already taken"))
	}
	s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	return handlerLogin(s, cmd)
}
