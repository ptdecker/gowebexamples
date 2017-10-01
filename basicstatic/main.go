/*
A basic static web site

Based upon Todd McLeod's tutorial at https://www.youtube.com/watch?v=joVuFbAzPYw
*/

package main

//TODO: Abstract repeated pattern in stagic page handlers into a single method
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
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/form", handleBasicForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, _ *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func handleAbout(w http.ResponseWriter, _ *http.Request) {
	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func handleContact(w http.ResponseWriter, _ *http.Request) {
	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
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

	err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
