package flag

import (
	"encoding/json"
	"flag"
	"testing"
)

func TestFlagsrc_Read(t *testing.T) {
	dbhost := flag.String("database-host", "", "db host")
	dbpw := flag.String("database-password", "", "db pw")

	flag.Set("database-host", "localhost")
	flag.Set("database-password", "some-password")
	flag.Parse()

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
	if actualDB["host"] != *dbhost {
		t.Errorf("expected %v got %v", *dbhost, actualDB["host"])
	}

	if actualDB["password"] != *dbpw {
		t.Errorf("expected %v got %v", *dbpw, actualDB["password"])
	}
}
