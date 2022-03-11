package process

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func UnmarshalYaml(name string) (data Summarize) {
	buf := loadFile(name)

	err := yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return
}

func loadFile(name string) []byte {
	if !isFileExist(name) {
		fmt.Print("error: this file does not exist\n")
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print("error: faild read " + name + "\n")
		os.Exit(1)
	}

	return buf
}
