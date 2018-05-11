package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type urlVal int

func (m urlVal) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
		URL         *url.URL
	}{
		req.Method,
		req.Form,
		req.URL,
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h urlVal
	http.ListenAndServe(":8080", h)
}
