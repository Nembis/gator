package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: login command requires 1 argument, username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error: you can't login to an account that doesn't exist!")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("error: failed to set username: %v", err)
	}

	fmt.Printf("Successfully set username to %s\n", s.cfg.CurrentUserName)

	return nil
}
