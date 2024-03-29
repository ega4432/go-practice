package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// https://golang.org/doc/articles/wiki/
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test
	title := r.URL.Path[len("/view/"):]
	title = "section13/" + title
	p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// "section13" という文字列を置換して edit 用のリンクを使いたい
	p.Title = strings.Replace(p.Title, "section13/", "", 1)
	renderTemplate(w, "section13/view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	title = "section13/" + title
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	p.Title = strings.Replace(p.Title, "section13/", "", 1)
	renderTemplate(w, "section13/edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	log.Println(title)
	p.Title = "section13/" + p.Title
	err := p.save()
	log.Println(title)
	p.Title = strings.Replace(p.Title, "section13/", "", 1)
	log.Println(title)
	if err != nil {
		log.Println("ERROR")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
