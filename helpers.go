package common

import (
	"encoding/json"
	"io/ioutil"
)

func LoadFile(path string, strct interface{}) (e error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &strct)
	if err != nil {
		return err
	}
	return
}
