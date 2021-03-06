package config

import (
	"testing"
)

func TestConfig_LogDir(t *testing.T) {
	cfg, err := DefaultConfig(DefaultDataDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cfg.Version)
	t.Log(cfg.DataDir)
}
