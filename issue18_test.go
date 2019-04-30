package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/goasana/config/source/env"
	"github.com/goasana/config/source/file"
)

func createFileForIssue18(t *testing.T, content string) *os.File {
	data := []byte(content)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	return fh
}

func TestIssue18(t *testing.T) {
	fh := createFileForIssue18(t, `{
  "amqp": {
    "host": "rabbit.platform",
    "port": 80
  },
  "handler": {
    "exchange": "springCloudBus"
  }
}`)
	path := fh.Name()
	defer func() {
		_ = fh.Close()
		_ = os.Remove(path)
	}()
	_ = os.Setenv("AMQP_HOST", "rabbit.testing.com")

	conf := NewConfig()
	_ = conf.Load(
		file.NewSource(
			file.WithPath(path),
		),
		env.NewSource(),
	)

	actualHost := conf.Get("amqp", "host").String("backup")
	if actualHost != "rabbit.testing.com" {
		t.Fatalf("Expected %v but got %v",
			"rabbit.testing.com",
			actualHost)
	}
}
