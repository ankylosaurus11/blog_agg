package main

import (
	"errors"
	"fmt"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
)

type state struct {
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
	handler, ok := c.Cmd[cmd.Name]; ok {
		err := handler(s, cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			return err
		} else {
			return fmt.Errorf("command not found: %s", cmd.Name)
		}
	return nil
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Command) == 0 {
		return errors.New("not enough arguments were provided")
	}
	s.ConfigPointer.CurrentUserName = cmd.Command[0]
	fmt.Println("set user:", cmd.Command[0])

	return nil
}
