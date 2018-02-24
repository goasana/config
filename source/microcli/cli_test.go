package microcli

import (
	"encoding/json"
	"github.com/micro/cli"
	"github.com/micro/go-config/source"
	"testing"
)

func TestClisrc_Read(t *testing.T) {
	var src source.Source
	app := cli.NewApp()
	app.Name = "testapp"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "db-host"},
	}
	app.Action = func(c *cli.Context) {
		src = NewSource(c)
	}
	app.Run([]string{"run", "-db-host", "localhost"})

	c, err := src.Read()
	if err != nil {
		t.Error(err)
	}

	var actual map[string]interface{}
	if err := json.Unmarshal(c.Data, &actual); err != nil {
		t.Error(err)
	}

	actualDB := actual["db"].(map[string]interface{})
	if actualDB["host"] != "localhost" {
		t.Errorf("expected localhost, got %s", actualDB["name"])
	}
}
