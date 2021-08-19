package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/goLocal_db_manag"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// city := r.FormValue("city")
	// if strings.Compare("", city) != 0 {

	// }
	p, err := loadPage("index")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "index", p)

	// Conecting to database
	goLocal_db_manag.ConnectDB()
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles((tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
    return nil, err
  }
	return &Page{Title: title, Body: body}, nil
}
