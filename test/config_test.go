package test

import (
	"github.com/Ruixi-rebirth/QuickMailer/pkg/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig("../config/test_config.json")
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	if len(cfg.FilePaths) != 2 {
		t.Errorf("Expected 2 file paths, got %d", len(cfg.FilePaths))
	}

	if cfg.SenderEmail != "test@example.com" {
		t.Errorf("Expected sender email to be 'test@example.com', got '%s'", cfg.SenderEmail)
	}

	if cfg.Subject != "Test Subject" {
		t.Errorf("Expected subject to be 'Test Subject', got '%s'", cfg.Subject)
	}
}
