package main

import (
	"fmt"

	"github.com/nembis/gator/internal/config"
	"github.com/nembis/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("error: command does not exist")
	}

	if err := f(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
