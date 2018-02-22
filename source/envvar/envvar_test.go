package envvar

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestEnvvar_Read(t *testing.T) {
	expected := map[string]map[string]string{
		"database": {
			"host":     "localhost",
			"password": "password",
		},
	}

	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PASSWORD", "password")

	source := NewSource()
	c, err := source.Read()
	if err != nil {
		t.Error(err)
	}

	var actual map[string]interface{}
	if err := json.Unmarshal(c.Data, &actual); err != nil {
		t.Error(err)
	}

	actualDB := actual["database"].(map[string]interface{})

	for k, v := range expected["database"] {
		a := actualDB[k]

		if a != v {
			t.Errorf("expected %v got %v", v, a)
		}
	}
}

func TestEnvvar_PrefixIgnoresOtherEnvs(t *testing.T) {
	os.Setenv("GOMICRO_DATABASE_HOST", "localhost")
	os.Setenv("GOMICRO_DATABASE_PASSWORD", "password")
	source := NewSource(WithPrefix("GOMICRO_"))

	c, err := source.Read()
	if err != nil {
		t.Error(err)
	}

	var actual map[string]interface{}
	if err := json.Unmarshal(c.Data, &actual); err != nil {
		t.Error(err)
	}

	if l := len(actual); l != 1 {
		t.Errorf("expected 1 top key, got %v", l)
	}
}

func TestEnvvar_WatchNextNoOpsUntilStop(t *testing.T) {
	source := NewSource(WithPrefix("GOMICRO_"))
	w, err := source.Watch()
	if err != nil {
		t.Error(err)
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		w.Stop()
	}()

	if _, err := w.Next(); err.Error() != "watcher stopped" {
		t.Errorf("expected watcher stopped error, got %v", err)
	}
}
