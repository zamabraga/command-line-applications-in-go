package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {

	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := parseContent(input)
	expect, err := ioutil.ReadFile(goldenFile)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expect, result) {
		t.Logf("golden:\n %s\n", expect)
		t.Logf("result:\n %s\n", result)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {

	var mockStOut bytes.Buffer

	if err := run(inputFile, &mockStOut, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStOut.String())
	result, err := ioutil.ReadFile(resultFile)

	if err != nil {
		t.Fatal(err)
	}

	expect, err := ioutil.ReadFile(goldenFile)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expect, result) {
		t.Logf("golden:\n %s\n", expect)
		t.Logf("result:\n %s\n", result)
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)

}
