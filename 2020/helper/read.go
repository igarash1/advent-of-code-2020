package helper

import (
	"io/ioutil"
	"log"
)

func ReadStringFromFile(fileName string) string {
	if inputs, err := ioutil.ReadFile(fileName); err != nil {
		log.Fatal(err)
	} else {
		return string(inputs)
	}
}
