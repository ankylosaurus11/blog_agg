package main

import gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"

type state struct {
	ConfigPointer *gatorconfig.Config
}

type command struct {
	Name     string
	Command []string
}

type commands struct {
	
}

func handlerLogin(s *state, cmd command) error {
	if cmd.Command = []string{} {
		errors.New("not enough arguments were provided")
	}
	s.ConfigPointer.Config.CurrentUserName = cmd.Command[1]
	fmt.SPrintf("user: %v, has been set.", cmd.Command[1])
}