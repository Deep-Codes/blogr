package utils

import (
	"fmt"
	"strconv"
)

// Toggle Read of a Blog
func Update(arg []string) {
	list := ReadJson("blogs.json")
	if len(arg) <= 2 {
		fmt.Println(`
		 Not Enough Arguments Passed use
		 -o <Index of Blog (Integer)>
		 `)
		return
	}
	if i, err := strconv.Atoi(arg[2]); err == nil {
		list[i-1].Read = !list[i-1].Read
		WriteJson(list)
		fmt.Println("Updated! ")
	} else {
		fmt.Println("Please Enter an Integer")
	}
}
