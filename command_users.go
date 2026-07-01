package main

import (
	"context"
	"fmt"
)

func handlersListUsers(s *state, cmd command) error {

	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: failed ot get all registered users: %v", err)
	}

	for _, user := range dbUsers {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}
