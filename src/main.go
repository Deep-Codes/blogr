package main

import (
	"fmt"
	"os"

	"../utils"
)

var helpLog = `
	Usage: 

	-l         lists all blogs
	-a <data>  adds blog
	-r <index> removes blog
	-u <index> updates blog
	-s <port>	 start server

`

func main() {
	arg := os.Args
	if len(arg) <= 1 {
		fmt.Print(helpLog)
		return
	}
	switch arg[1] {
	case "-l":
		utils.List()
	case "-a":
		utils.Add()
	case "-r":
		utils.Remove(arg)
	case "-u":
		utils.Update(arg)
	case "-o":
		utils.Open(arg)
	case "-s":
		utils.StartServer()
	default:
		fmt.Println("Invalid Flag , Use the Following : ")
		fmt.Println(helpLog)
	}
}
