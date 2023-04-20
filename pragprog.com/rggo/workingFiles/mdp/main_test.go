package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "./test1.md.html"
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

	if err := run(inputFile); err != nil {
		t.Fatal(err)
	}

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
