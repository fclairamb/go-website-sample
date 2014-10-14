package main

import (
	"html/template"
	"net/http"
)

var TEMPLATE_LOGIN = template.Must(template.ParseFiles("templates/_main.html", "templates/login.html"))

type PageLoginData struct {
	PageBaseData
	JustAuthenticated bool
	Nb                int
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	data := &PageLoginData{}
	data.Title = "Login"
	data.Session = session

	if !data.Session.Authenticated {
		if r.Method == "POST" {
			username := r.PostFormValue("username")
			password := r.PostFormValue("password")
			data.Session.Authenticate(username, password)
			data.JustAuthenticated = session.Authenticated
			if data.Session.Authenticated {
				if r.PostFormValue("remember") == "true" {
					log.Debug("Authenticated with remember !")
					data.Session.SaveCookies(w, r)
				}

				if source := r.URL.Query().Get("source"); source != "" {
					http.Redirect(w, r, source, http.StatusFound)
				}
			}
		}
	}

	TEMPLATE_LOGIN.Execute(w, data)
}
