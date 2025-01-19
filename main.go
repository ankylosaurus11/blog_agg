package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

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

	commands.register("register", handlerRegister)
	commands.register("login", handlerLogin)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("feeds", handlerFeeds)
	commands.register("following", following)
	commands.register("agg", func(s *state, cmd command) error {
		if len(cmd.Command) != 1 {
			return errors.New("agg requires a single duration")
		}
		timeRequested := cmd.Command[0]
		timeBetweenRequests, err := time.ParseDuration(timeRequested)
		if err != nil {
			return err
		}
		ticker := time.NewTicker(timeBetweenRequests)
		fmt.Println("Collecting feeds every ", timeRequested)
		for ; ; <-ticker.C {
			scrapeFeeds(s)
		}
	})
	commands.register("addfeed", middlewareLoggedIn(func(s *state, cmd command, user database.User) error {
		if len(cmd.Command) < 2 {
			return errors.New("addfeed requires two arguments: name and url")
		}
		name := cmd.Command[0]
		url := cmd.Command[1]
		return addFeed(s, cmd, name, url)
	}))
	commands.register("follow", middlewareLoggedIn(func(s *state, cmd command, user database.User) error {
		if len(cmd.Command) < 1 {
			return errors.New("follow requires an argument: url")
		}
		url := cmd.Command[0]
		return follow(s, cmd, url)
	}))
	commands.register("unfollow", middlewareLoggedIn(func(s *state, cmd command, user database.User) error {
		if len(cmd.Command) < 1 {
			return errors.New("unfollow requires an argument: url")
		}
		url := cmd.Command[0]
		return unfollow(s, cmd, url)
	}))

	if len(os.Args) < 2 {
		fmt.Println("Not enough commands, please provide at least a command name")
	}
	commandName := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		Name:    commandName,
		Command: args,
	}

	if err := commands.run(&newState, cmd); err != nil {
		log.Fatal(err)
	}
}
