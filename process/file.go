package process

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadYamlFile(name string) (data YamlFile) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print("error: faild read " + name + "\n")
		os.Exit(1)
	}

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return
}
