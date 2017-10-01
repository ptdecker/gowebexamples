/*
A basic static web site

Based upon Todd McLeod's tutorial at https://www.youtube.com/watch?v=joVuFbAzPYw
*/

package main

//TODO: Extend pageData structure in here instead of adding name to global page data

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", handleStatic("Index", "index.html"))
	http.HandleFunc("/about", handleStatic("About", "about.html"))
	http.HandleFunc("/contact", handleStatic("Contact", "contact.html"))
	http.HandleFunc("/form", handleBasicForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// Middleware wrapper component to abstract common behaviors for static pages served with a template

func handleStatic(pageTitle string, pageURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		pd := pageData{
			Title: pageTitle,
		}
		err := tpl.ExecuteTemplate(w, pageURL, pd)
		if err != nil {
			log.Println("LOGGED", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func handleBasicForm(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Basic Form",
	}

	var first string

	if req.Method == http.MethodPost {
		first = req.FormValue("firstName")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
