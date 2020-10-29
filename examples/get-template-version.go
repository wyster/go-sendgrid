package main

import (
	"fmt"
	"os"
	"sendgrid/template"
)

func main() {
	responseData := template.GetVersion(
		os.Getenv("SENDGRID_TOKEN"),
		os.Getenv("SENDGRID_TEMPLATE_ID"),
		os.Getenv("SENDGRID_TEMPLATE_VERSION"),
	)

	fmt.Printf("Response: %+v", responseData)
}
