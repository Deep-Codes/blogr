package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func WriteJson(data []Blog) {
	b, _ := json.Marshal(data)
	ioutil.WriteFile("blogs.json", b, os.ModePerm)
}
