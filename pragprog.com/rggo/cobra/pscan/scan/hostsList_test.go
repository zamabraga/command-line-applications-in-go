package scan_test

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"pragpro.com/rggo/cobra/pscan/scan"
)

func TestAdd(t *testing.T) {
	testCase := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"AddNew", "hosts2", 2, nil},
		{"AddExists", "hosts1", 1, scan.ErrExists},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}
			if err := hl.Add("hosts1"); err != nil {
				t.Fatal(err)
			}

			err := hl.Add(tc.host)
			if tc.expectErr != nil {
				if err == nil {
					t.Fatal("Expected error, got nil instead\n")
				}

				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error %q, got %q instead\n", tc.expectErr, err)

				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}

			if hl.Hosts[1] != tc.host {
				t.Errorf("Expected host name %q, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}
func TestRemove(t *testing.T) {
	testCase := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"RemoveExisting", "hosts1", 1, nil},
		{"RemoveNotFound", "hosts2", 1, scan.ErrNotExists},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			for _, h := range []string{"hosts1", "hosts2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}
			err := hl.Remove(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatal("Expected error, got nil instead\n")
				}

				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error %q, got %q instead\n", tc.expectErr, err)

				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}

			if hl.Hosts[0] != tc.host {
				t.Errorf("Expected host name %q shoud not be in list\n", tc.host)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hl1 := scan.HostsList{}
	hl2 := scan.HostsList{}

	hostName := "host1"
	hl1.Add(hostName)

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("Host %q should match %q host", hl1.Hosts[0], hl2.Hosts[0])
	}
}
func TestLoadNoFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error deleting temp file: %s", err)
	}

	hl := &scan.HostsList{}

	if err := hl.Load(tf.Name()); err != nil {
		t.Errorf("Expected no error, got %q instead\n", err)
	}

}
