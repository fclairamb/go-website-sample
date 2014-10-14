package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerCore)
	http.HandleFunc("/sample2", handlerSample2)
	http.HandleFunc("/about", handlerAbout)
	http.HandleFunc("/home", handlerHome)
	http.HandleFunc("/session", handlerSession)
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/logout", handlerLogout)
	http.HandleFunc("/complete", handlerComplete)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":8787", nil)
}
