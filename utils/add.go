package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')
	fmt.Print("Enter URL: ")
	url, _ := reader.ReadString('\n')
	list := ReadJson("blogs.json")
	list = append(list, Blog{
		Title: strings.Trim(title, "\n"),
		Url:   strings.Trim(url, "\n"),
		Read:  false,
	})
	WriteJson(list)
	fmt.Println("Added! ")
}
