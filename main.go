package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.html")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(PORT, nil)
}
