package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/pasien/index.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/pasien/add.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/pasien/edit.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/pasien/delete.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
