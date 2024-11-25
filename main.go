package main

import (
	"fmt"
	"log"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
)

func main() {
	cfg, err := gatorconfig.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	cfg.SetUser("Dylan")

	updatedCfg, err := gatorconfig.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	fmt.Println(updatedCfg)
}
