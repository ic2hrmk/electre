package main

import (
	"io/ioutil"
	"encoding/json"
	"electre/electre"
)



func readFile(fileName string) (variant *electre.Variant, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &variant)
	return
}