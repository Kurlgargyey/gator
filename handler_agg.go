package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: agg <time-between-reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("time between reqs was not a valid duration string: %s. error: %w", cmd.args[0], err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
