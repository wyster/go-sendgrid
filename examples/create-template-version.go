package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
	"log"
	"os"
)

func main() {
	responseData, err := template.CreateVersion(
		os.Getenv("SENDGRID_TOKEN"),
		os.Getenv("SENDGRID_TEMPLATE_ID"),
		"test",
	)
	if err != nil {
		log.Fatal("Template version not created! ", err)
	}

	fmt.Printf("Response: %+v", responseData)
}
