package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/processor", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// get form values
	firster := r.FormValue("firster")
	laster := r.FormValue("laster")

	// process form values
	d := struct {
		Firster string
		Laster  string
	}{
		Firster: firster,
		Laster:  laster,
	}

	// render template
	tpl.ExecuteTemplate(w, "processor.gohtml", d)

}
