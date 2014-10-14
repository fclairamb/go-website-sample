package main

import (
	"html/template"
	"net/http"
)

var TEMPLATE_COMPLETE = template.Must(template.ParseFiles("templates/_main.html", "templates/test.html"))

type PageTestData struct {
	PageBaseData
}

func handlerComplete(w http.ResponseWriter, r *http.Request) {
	data := &PageTestData{}
	data.Title = "About"
	data.Session = GetSession(w, r)

	TEMPLATE_COMPLETE.Execute(w, data)
}
