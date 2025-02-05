package main

import "gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	dict map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.dict[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	return c.dict[cmd.name](s, cmd)
}