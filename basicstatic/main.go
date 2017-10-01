/*
A basic static web site

Based upon Todd Mcleod's tutorial at https://www.youtube.com/watch?v=joVuFbAzPYw
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

// GlobalPageData structure holds variables passed to all pages

type GlobalPageData struct {
	Title     string
}

// Global tpl contains pointer to all parsed templates

var tpl *template.Template

// Get things going by parsing templates

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Set up all routes to static pages and serve things up

func main() {
	http.HandleFunc("/", handleStatic("Index", "index.html"))
	http.HandleFunc("/about", handleStatic("About", "about.html"))
	http.HandleFunc("/contact", handleStatic("Contact", "contact.html"))
	http.HandleFunc("/form", handleBasicForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./pub"))))
	http.ListenAndServe(":8080", nil)
}

// Middleware wrapper component to abstract common behaviors for static pages served with a template

func handleStatic(pageTitle string, pageURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		pd := GlobalPageData{
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

// A basic web form handler

func handleBasicForm(w http.ResponseWriter, req *http.Request) {

	type LocalPageData struct {
		GlobalPageData
		FirstName string
	}

	var pd LocalPageData
	pd.Title = "Basic Form"
	if req.Method == http.MethodPost {
		pd.FirstName = req.FormValue("firstName")
	}

	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
