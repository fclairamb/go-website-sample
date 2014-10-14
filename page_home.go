package main

import (
	"html/template"
	"net/http"
)

var TEMPLATE_HOME = template.Must(template.ParseFiles("templates/_main.html", "templates/home.html"))

type PageHomeData struct {
	PageBaseData
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	data := &PageHomeData{}
	data.Title = "Home"
	data.Session = GetSession(w, r)

	TEMPLATE_HOME.Execute(w, data)
}
