package pefiles

import (
	"io/ioutil"
	"testing"
)

func TestCheckMZSignature(t *testing.T) {
	testBytes, err := ioutil.ReadFile("../../testfiles/test.exe")
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
