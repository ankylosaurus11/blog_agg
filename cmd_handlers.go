package main

import (
	"errors"
	"fmt"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
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
