package main

import gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"

type state struct {
	ConfigPointer *gatorconfig.Config
}

type command struct {
	Name     string
	Commands []string
}
