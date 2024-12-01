package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
)

func main() {
	cfg, err := gatorconfig.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var newState state

	newState.ConfigPointer = &cfg

	db, err := sql.Open("postgres", cfg.DBURL)

	dbQueries := database.New(db)

	commands := commands{
		Cmd: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Not enough commands, please provide at least a command name")
	}
	commandName := os.Args[1]
	args := os.Args[2:]

	if len(args) == 0 {
		fmt.Println("username is required")
	}

	cmd := command{
		Name:    commandName,
		Command: args,
	}

	if err := commands.run(&newState, cmd); err != nil {
		log.Fatal(err)
	}
}
