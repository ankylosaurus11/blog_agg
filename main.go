package main

import (
	"context"
	"database/sql"
	"errors"
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

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	var newState state
	newState.ConfigPointer = &cfg
	newState.db = database.New(db)

	commands := commands{
		Cmd: make(map[string]func(*state, command) error),
	}

	ctx := context.Background()

	commands.register("register", handlerRegister)
	commands.register("login", handlerLogin)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("feeds", handlerFeeds)
	commands.register("agg", func(s *state, cmd command) error {
		rssFeed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
		if err != nil {
			return err
		}
		fmt.Println(rssFeed)
		return nil
	})
	commands.register("addfeed", func(s *state, cmd command) error {
		if len(cmd.Command) < 2 {
			return errors.New("addfeed requires two arguments: name and url")
		}
		name := cmd.Command[0]
		url := cmd.Command[1]
		return addFeed(s, cmd, name, url)
	})

	if len(os.Args) < 2 {
		fmt.Println("Not enough commands, please provide at least a command name")
	}
	commandName := os.Args[1]
	args := os.Args[2:]

	if len(args) == 0 && commandName != "reset" && commandName != "users" && commandName != "agg" && commandName != "addfeed" && commandName != "feeds" {
		log.Fatal("username is required")
	}

	cmd := command{
		Name:    commandName,
		Command: args,
	}

	if err := commands.run(&newState, cmd); err != nil {
		log.Fatal(err)
	}
}
