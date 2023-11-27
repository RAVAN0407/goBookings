package handler

import (
	"html/template"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("base.html.tmpl").ParseFiles("/home/bhanu/projects/goBookings/pkg/templates/base.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
