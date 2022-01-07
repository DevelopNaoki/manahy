package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
)

func ReadFile(name string) (data YamlFile) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print("error: faild read "+name+"\n")
		os.Exit(1)
	}

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return
}
