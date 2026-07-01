package main

import (
	"context"
	"fmt"

	"github.com/nembis/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		dbUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error: failed to get user from database: %v", err)
		}

		return handler(s, c, dbUser)
	}
}
