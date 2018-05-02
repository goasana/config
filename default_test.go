package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/micro/go-config/source/file"
)

func TestLoadWithGoodFile(t *testing.T) {
	data := []byte(`{"foo": "bar"}`)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		fh.Close()
		os.Remove(path)
	}()
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	// Create new config
	conf := NewConfig()
	// Load file source
	if err := conf.Load(file.NewSource(
		file.WithPath(path),
	)); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
}

func TestLoadWithInvalidFile(t *testing.T) {
	// Create new config
	conf := NewConfig()
	// Load file source
	if err := conf.Load(file.NewSource(
		file.WithPath("/i/do/not/exists.json"),
	)); err == nil {
		t.Fatal("Expected error but none !")
	}
}
