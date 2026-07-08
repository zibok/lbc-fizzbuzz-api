package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/zibok/lbc-fizzbuzz-api/internal/httpapi"
)

func loadConfig(path string) (httpapi.Config, error) {
	config := httpapi.DefaultConfig()
	if path == "" {
		return config, nil
	}

	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return config, nil
		}
		return httpapi.Config{}, fmt.Errorf("open config file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return httpapi.Config{}, fmt.Errorf("decode config file: %w", err)
	}

	config = config.WithDefaults()
	if config.MaxLimit < 1 {
		return httpapi.Config{}, fmt.Errorf("maxLimit must be greater than 0")
	}

	return config, nil
}
