package utils

import (
	"fmt"
)

func isRead(b bool) string {
	if b {
		return "X"
	} else {
		return " "
	}
}

func formatStr(s Blog, i int) string {
	str := fmt.Sprintf("%d. [%s] %s ", i+1, isRead(s.Read), s.Title)
	return str
}

func List() {
	list := ReadJson("blogs.json")
	fmt.Println("")
	fmt.Println("All Blogs:")
	fmt.Println("")
	for i := 0; i < len(list); i++ {
		fmt.Println(formatStr(list[i], i))
	}
	fmt.Println("")
}
