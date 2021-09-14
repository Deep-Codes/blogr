package utils

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"
)

type PageVariables struct {
	Blogs []Blog
}

func launch(url string) {
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

func getBlogs(w http.ResponseWriter, r *http.Request) {
	list := ReadJson("blogs.json")
	pageVariables := PageVariables{
		Blogs: list,
	}
	t, err := template.ParseFiles("todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template Parsing Error: ", err)
	}
	t.Execute(w, pageVariables)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	list := ReadJson("blogs.json")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request Parsing Error: ", err)
	}
	blog := Blog{
		Title: r.FormValue("Title"),
		Url:   r.FormValue("Url"),
		Read:  false,
	}
	list = append(list, blog)
	WriteJson(list)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func StartServer() {
	launch("http://localhost:8080/")
	http.HandleFunc("/", getBlogs)
	http.HandleFunc("/add-blog", addBlog)
	fmt.Println("Server is running on PORT : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
