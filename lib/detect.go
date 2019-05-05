package lib

import (
	"io/ioutil"
)

func Detect(p string) string {

	files, err := ioutil.ReadDir(p)

	if err != nil {

	}
	result := ""
	for _, file := range files {

		if file.IsDir() {
			continue
		}
		n := file.Name()
		detect(n, &result)
		if result != "" {
			break
		}
	}

	return result

}

func detect(file string, result *string) {

	if file == "setup.py" {
		*result = "python"
	}
	if file == "go.mod" {
		*result = "go"
	}

}
