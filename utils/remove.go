package utils

import (
	"fmt"
	"strconv"
)

func removeIndex(s []Blog, index int) []Blog {
	return append(s[:index], s[index+1:]...)
}

// Remove the Blog
func Remove(arg []string) {
	list := ReadJson("blogs.json")
	if len(arg) <= 2 {
		fmt.Println(`
		 Not Enough Arguments Passed use
		 -r <Index of Blog (Integer)>
		 `)
		return
	}
	if i, err := strconv.Atoi(arg[2]); err == nil {
		newList := removeIndex(list, i-1)
		WriteJson(newList)
		fmt.Println("Removed! ")
	} else {
		fmt.Println("Please Enter an Integer")
	}
}
