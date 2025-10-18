package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: reset expects zero arguments")
	}

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("error: deleting users: %w", err)
	}

	fmt.Println("users were successfully reset")
	return nil
}
