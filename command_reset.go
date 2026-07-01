package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("error: failed to reset database: %v", err)
	}

	fmt.Println("Successfully reset database!")

	return nil
}
