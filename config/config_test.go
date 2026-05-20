package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("../config.yaml")
	if err != nil {
		t.Errorf("LoadConfig failed: %v", err)
	}
}
