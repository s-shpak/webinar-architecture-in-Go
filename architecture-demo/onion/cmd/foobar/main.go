package main

import (
	"fmt"
	"log"

	"onion/internal/core/application"
	"onion/internal/infra/api/rest"
	"onion/internal/infra/store"
	"onion/internal/infra/store/memory"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	store, err := store.NewStore(store.Config{
		Memory: &memory.Config{},
	})
	if err != nil {
		return fmt.Errorf("failed to initialize a store: %w", err)
	}
	foobar := application.NewApplication(store)
	api := rest.NewAPI(foobar)
	if err := api.Run(); err != nil {
		return fmt.Errorf("server has failed: %w", err)
	}
	return nil
}
