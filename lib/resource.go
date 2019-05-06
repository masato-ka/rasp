package lib

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Getrepo(path string, dst string) (string, error) {

	abs, err := filepath.Abs(dst)
	folder := GetFolder(path)
	fullpath := filepath.Join(abs, folder)

	_, err = os.Stat(fullpath)

	var args []string
	if os.IsNotExist(err) {
		args = append(args, "clone", path)

	} else {
		args = append(args, "pull")
		dst = fullpath
	}

	//buffer := &bytes.Buffer{}
	ps, err := Execution("git", args, dst, os.Stdout, os.Stdout)
	if err != nil {

		fmt.Fprintf(os.Stderr, "Failed run git command %+v \n", err)
		return "", err
	}
	//output := buffer.String()
	//folder := strings.Split(output, "'")[1]

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed get dst abs path.: %+v \n", err)
		return "", err
	}
	fmt.Fprintf(os.Stdout, "%s\n", ps.String())
	return fullpath, nil
}

func GetFolder(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		return ""
	}
	parsed := strings.Split(u.Path, "/")
	folder := strings.Split(parsed[2], ".")[0]

	return folder
}
