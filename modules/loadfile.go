package modules

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func UnmarshalYaml(name string) (data Summarize, err error) {
	buf, err := loadFile(name)
	if err != nil {
		return data, err
	}

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func loadFile(name string) (buf []byte, err error) {
	fileExist, err := isFileExist(name)
	if err != nil {
		return nil, err
	} else if !fileExist {
		return nil, fmt.Errorf("error: this file does not exist\n")
	}

	buf, err = ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("error: faild read %s\n", name)
	}
	return buf, nil
}
