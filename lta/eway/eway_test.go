package eway

import (
	"io"
	"os"
	"testing"
)

func testRdr(t *testing.T) io.Reader {
	f, err := os.Open("testdata/output.json")
	if err != nil {
		t.Fatal(err)
	}

	return f
}

func testData(t *testing.T) estTravelTime {
	return load(testRdr(t))
}

func TestLoad(t *testing.T) {
	eTime := testData(t)
	t.Logf("%s", eTime)
}

