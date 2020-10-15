package pefiles

import (
	"io/ioutil"
	"testing"
)

func TestCheckMZSignature(t *testing.T) {
	testfiles, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatal("Could not get testfiles")
	}
	for _, testfile := range testfiles {
		testfileName := "testdata/" + testfile.Name()
		testBytes, err := ioutil.ReadFile(testfileName)
		if err != nil {
			t.Fatal("Could not open testfile")
		}
		MZ, err := checkMzSignature(testBytes)
		if err != nil {
			t.Fatal("Could not open testfile")
		}
		if MZ != true {
			t.Fatal("Failed to verify MZ singature")
		}
	}
}
