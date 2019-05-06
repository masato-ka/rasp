package lib

import (
	"rasp/lib"
	"testing"
)

func TestGetFolderNormal01(t *testing.T) {
	folder := lib.GetFolder("https://github.com/masato-ka/rasp")

	if folder != "rasp" {
		t.Fail()
	}

}

func TestGetFolderNormal02(t *testing.T) {
	folder := lib.GetFolder("https://github.com/masato-ka/rasp.git")

	if folder != "rasp" {
		t.Fail()
	}

}

func TestGetFolderNormal03(t *testing.T) {
	folder := lib.GetFolder("https://github.com/masato-ka/platform-wiolte/tree/bug-fix_change-uploadtool")

	if folder != "platform-wiolte" {
		t.Fail()
	}

}
