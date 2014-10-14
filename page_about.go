package main

import (
	"html/template"
	"net/http"
)

var TEMPLATE_ABOUT = template.Must(template.ParseFiles("templates/_main.html", "templates/about.html"))

type PageAboutData struct {
	PageBaseData
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	data := &PageAboutData{}
	data.Title = "About"
	data.Session = GetSession(w, r)

	TEMPLATE_ABOUT.Execute(w, data)
}
