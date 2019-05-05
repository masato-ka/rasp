package lib

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Getrepo(path string, dst string) (string, error) {
	args := []string{"clone", path}
	buffer := &bytes.Buffer{}
	ps, err := Execution("git", args, dst, os.Stdout, buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed run git command %+v \n", err)
		return "", err
	}
	output := buffer.String()
	folder := strings.Split(output, "'")[1]

	abs, err := filepath.Abs(dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed get dst abs path.: %+v \n", err)
		return "", err
	}
	fullpath := filepath.Join(abs, folder)
	fmt.Fprintf(os.Stdout, "%s\n", ps.String())
	return fullpath, nil
}
