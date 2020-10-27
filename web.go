package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sendgrid/template"
)

func templatesListHandler(w http.ResponseWriter, r *http.Request) {
	responseData := template.List(os.Getenv("SENDGRID_TOKEN"))
	fmt.Fprintln(w, "<html><ul>")
	for _, id := range responseData.IdsList() {
		fmt.Fprintf(w, "<li><a href='/template/show/%s'>%s</a></li>", id, id)
	}
	fmt.Fprintln(w, "</ul></html>")
}

func templateShowHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/template/show/"):]
	responseData := template.Get(os.Getenv("SENDGRID_TOKEN"), id)
	if len(responseData.Versions) > 0 {
		fmt.Fprintf(w, "%s", responseData.Versions[0].HtmlContent)
	} else {
		fmt.Fprintf(w, "%s", "template versions not found")
	}
}

func main() {
	http.HandleFunc("/template/list/", templatesListHandler)
	http.HandleFunc("/template/show/", templateShowHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
