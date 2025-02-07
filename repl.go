package main

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if cm, ok := c.handlers[cmd.name]; ok {
		return cm(s, cmd)
	} else {
		return fmt.Errorf("unknown command")
	}
}
