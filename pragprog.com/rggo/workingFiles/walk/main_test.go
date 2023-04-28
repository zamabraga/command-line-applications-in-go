package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		root   string
		cfg    config
		expect string
	}{
		{"NoFilter", "testdata", config{ext: "", size: 0, list: true}, "testdata/dir.log\ntestdata/dir2/script.sh\n"},
		{"FilterExtensionSizeNoMatch", "testdata", config{ext: ".log", size: 10, list: true}, "testdata/dir.log\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expect != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expect, res)
			}
		})
	}
}
