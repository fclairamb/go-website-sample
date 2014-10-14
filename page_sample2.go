package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Sample2Data struct {
	PageBaseData
	Persons []*Person
}

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
	Year     int
}

var TEMPLATE_SAMPLE2 = template.Must(template.ParseFiles("templates/_main.html", "templates/sample2.html"))

func handlerSample2(w http.ResponseWriter, r *http.Request) {
	data := &Sample2Data{
		Persons: []*Person{
			&Person{
				Name:   "Jan",
				Age:    50,
				Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
				Jobs: []*Job{
					&Job{Employer: "Monash", Role: "Honorary", Year: 2010},
					&Job{Employer: "Box Hill", Role: "Head of HE", Year: 2013},
					&Job{Employer: "Box Hill", Role: "Head of HE2", Year: 2014},
				},
			},
			&Person{
				Name:   "Florent",
				Age:    28,
				Emails: []string{"florent@clairambault.fr"},
			},
		},
	}
	data.Title = "Persons"
	data.Session = GetSession(w, r)

	TEMPLATE_SAMPLE2.Execute(w, data)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
