package file

import (
	"testing"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		p string
		f string
	}{
		{"/foo/bar.json", "json"},
		{"/foo/bar.yaml", "yaml"},
		{"/foo/bar.xml", "xml"},
		{"/foo/bar.conf.ini", "ini"},
		{"conf", "json"},
	}

	for _, d := range testCases {
		f := format(d.p)
		if f != d.f {
			t.Fatalf("%s: expected %s got %s", d.p, d.f, f)
		}
	}

}
