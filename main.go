package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("ravi")
	cfgNew := config.Read()
	fmt.Println(*cfgNew)
}