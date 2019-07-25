package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMustLoadConfig(t *testing.T) {
	cfg := MustLoadConfig()
	if os.Getenv("BASIC_ENV") != "prod" && fmt.Sprintf("%v:%v", cfg.Host, cfg.Port) != "localhost:8080" {
		t.Error("Not serving app ")
	}
}
