package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Opens Urls via Browsers
func Open(arg []string) {
	list := ReadJson("blogs.json")
	if len(arg) <= 2 {
		fmt.Println(`
		 Not Enough Arguments Passed use
		 -o <Index of Blog (Integer)>
		 `)
		return
	}
	if i, err := strconv.Atoi(arg[2]); err == nil {
		openBrowser(list[i-1].Url)
	} else {
		fmt.Println("Please Enter an Integer")
	}
}
