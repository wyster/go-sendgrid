package main

import (
	"fmt"
	"go-sendgrid/template"
	"log"
	"os"
)

func main() {
	responseData, err := template.GetVersion(
		os.Getenv("SENDGRID_TOKEN"),
		os.Getenv("SENDGRID_TEMPLATE_ID"),
		os.Getenv("SENDGRID_TEMPLATE_VERSION"),
	)
	if err != nil {
		log.Fatal("Templates not fetched! ", err)
	}

	fmt.Printf("Response: %+v", responseData)
}
