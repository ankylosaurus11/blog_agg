package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db            *database.Queries
	ConfigPointer *gatorconfig.Config
}

type command struct {
	Name    string
	Command []string
}

type commands struct {
	Cmd map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Cmd[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.Cmd[cmd.Name]
	if ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("command not found: %s", cmd.Name)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Command) == 0 {
		return errors.New("not enough arguments were provided")
	}
	if err := s.ConfigPointer.SetUser(cmd.Command[0]); err != nil {
		return err
	}
	fmt.Println("set user:", cmd.Command[0])

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Command) == 0 {
		return errors.New("not enough arguments were provided")
	}

	ctx := context.Background()
	userID := uuid.New()
	now := time.Now()

	_, err := s.db.GetUser(ctx, cmd.Command[0])
	if err == nil {
		return errors.New("user already exists")
	}

	var newUser database.User
	newUser, err = s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        userID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.Command[0],
	})
	if err != nil {
		return err
	}

	if err := s.ConfigPointer.SetUser(newUser.Name); err != nil {
		return err
	}
	fmt.Println("Created user: ", newUser.Name)
	return nil
}
