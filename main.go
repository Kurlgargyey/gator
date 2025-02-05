package main

import (
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	cfg := config.Read()
	s := state{
		cfg,
	}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("You did not provide enough arguments!")
	}
	args = args[1:]
	err := cmds.run(&s, command{
		name: args[0],
		args: args[1:],
	})
	if err != nil {
		log.Fatal(err)
	}
}