package main

import (
	"fmt"
	"log"
	"os"
	"sendgrid/template"
)

func main() {
	deleted, err := template.Delete(os.Getenv("SENDGRID_TOKEN"), os.Getenv("SENDGRID_TEMPLATE_ID"))
	if err != nil {
		log.Fatal("Template not deleted! ", err)
	}
	if deleted {
		fmt.Println("Created!")
	}
}
