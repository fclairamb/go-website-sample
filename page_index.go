package main

import (
	"html/template"
	"net/http"
)

type PageNotFoundData struct {
	PageBaseData
}

var TEMPLATE_NOTFOUND = template.Must(template.ParseFiles("templates/_main.html", "templates/notfound.html"))

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	data := &PageNotFoundData{}
	data.Title = "Not found"
	data.Session = GetSession(w, r)
	TEMPLATE_NOTFOUND.Execute(w, data)
}

func handlerCore(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/" {
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		handlerNotFound(w, r)
	}
}
