package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	data := []byte(`{"addr":":9090","maxLimit":42}`)

	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	got, err := loadConfig(path)
	if err != nil {
		t.Fatalf("loadConfig() error = %v", err)
	}

	if got.Addr != ":9090" {
		t.Fatalf("Addr = %q, want %q", got.Addr, ":9090")
	}

	if got.MaxLimit != 42 {
		t.Fatalf("MaxLimit = %d, want 42", got.MaxLimit)
	}
}

func TestLoadConfigRejectsInvalidMaxLimit(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	data := []byte(`{"maxLimit":-1}`)

	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	if _, err := loadConfig(path); err == nil {
		t.Fatal("loadConfig() error = nil, want error")
	}
}
