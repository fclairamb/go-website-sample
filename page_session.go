package main

import (
	//"fmt"
	//besession "github.com/astaxie/beego/session"
	"html/template"
	"net/http"
)

var TEMPLATE_SESSTEST = template.Must(template.ParseFiles("templates/_main.html", "templates/session.html"))

type PageSessTestData struct {
	PageBaseData
	Nb int
}

func handlerSession(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)

	if !session.AuthenticatedIfNotRedirect(w, r) {
		return
	}

	data := &PageSessTestData{}
	data.Title = "Session test"
	data.Session = session

	if session.Get("nb") == nil {
		session.Set("nb", 0)
	} else {
		session.Set("nb", session.Get("nb").(int)+1)
	}

	data.Nb = session.Get("nb").(int)

	TEMPLATE_SESSTEST.Execute(w, data)
}
