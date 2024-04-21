package main

import (
	"net/http"

	"github.com/Muhless/GO-CRUD-VALIDATION/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/pasien", controllers.Index)
	http.HandleFunc("/pasien/index", controllers.Index)
	http.HandleFunc("/pasien/add", controllers.Add)
	http.HandleFunc("/pasien/edit", controllers.Edit)
	http.HandleFunc("/pasien/delete", controllers.Delete)

	http.ListenAndServe(":9000", nil)
}