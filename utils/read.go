package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadJson(file string) []Blog {
	filename, err := os.Open("blogs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	var result []Blog

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result
}
