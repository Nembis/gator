package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nembis/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: register command requires 1 argument, username")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("error: failed to generate uuid: %v", err)
	}

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return fmt.Errorf("error: failed to create user %v", err)
	}

	if err = s.cfg.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("error: failed to set username: %v", err)
	}

	fmt.Printf("Successfully registered user %s\n", cmd.args[0])
	fmt.Printf("User: %v\n", dbUser)

	return nil
}
