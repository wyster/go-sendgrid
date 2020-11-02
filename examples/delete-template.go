package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
	"log"
	"os"
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
