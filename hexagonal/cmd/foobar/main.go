package main

import (
	"fmt"
	"log"

	"hexagonal/internal/adapters/api/rest"
	"hexagonal/internal/adapters/store"
	"hexagonal/internal/adapters/store/memory"
	"hexagonal/internal/config"
	"hexagonal/internal/core/services/foobar"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	store, err := store.NewStore(config.Config{
		Store: store.Config{
			Memory: &memory.Config{},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to initialize a store: %w", err)
	}
	foobar := foobar.NewFoobar(store)
	api := rest.NewAPI(foobar)
	if err := api.Run(); err != nil {
		return fmt.Errorf("server has failed: %w", err)
	}
	return nil
}
