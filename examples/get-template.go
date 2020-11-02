package main

import (
	"fmt"
	"go-sendgrid/template"
	"log"
	"os"
)

func main() {
	responseData, err := template.Get(os.Getenv("SENDGRID_TOKEN"), os.Getenv("SENDGRID_TEMPLATE_ID"))
	if err != nil {
		log.Fatal("Template not fetched! ", err)
	}
	fmt.Printf("Response: %+v", responseData)
}
