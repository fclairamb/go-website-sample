package main

import (
	"html/template"
	"net/http"
)

var TEMPLATE_LOGOUT = template.Must(template.ParseFiles("templates/_main.html", "templates/logout.html"))

type PageLogoutData struct {
	PageBaseData
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	data := &PageLogoutData{}
	data.Title = "Logout"
	data.Session = session

	session.End(w, r)

	TEMPLATE_LOGOUT.Execute(w, data)
}
