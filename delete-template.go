package main

import (
	"fmt"
	"log"
	"sendgrid/template"
)

func main() {
	deleted, err := template.Delete()
	if err != nil {
		log.Fatal("Template not deleted! ", err)
	}
	if deleted {
		fmt.Println("Deleted!")
	}
}
