package lib

import (
	"rasp/lib"
	"testing"
)

func TestDetectNormal01(t *testing.T) {

	d := lib.Detect("../data/goproj")
	if d != "go" {
		t.Fail()
	}

}

func TestDetectNormal02(t *testing.T) {

	d := lib.Detect("../data/pythonproj")
	if d != "python" {
		t.Fail()
	}

}

func TestDetectNormal03(t *testing.T) {
	d := lib.Detect("../data/normalproj")
	if d != "" {
		t.Fail()
	}
}
