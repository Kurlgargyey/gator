package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Read()
	dbConn, err := sql.Open("postgres", cfg.DbURL)
	db := database.New(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	s := state{
		cfg,
		db,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("You did not provide enough arguments!")
	}
	args = args[1:]

	err = cmds.run(&s, command{
		name: args[0],
		args: args[1:],
	})
	if err != nil {
		log.Fatal(err)
	}
}
