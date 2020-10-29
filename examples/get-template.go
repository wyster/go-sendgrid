package main

import (
	"fmt"
	"os"
	"sendgrid/template"
)

func main() {
	responseData := template.Get(os.Getenv("SENDGRID_TOKEN"), os.Getenv("SENDGRID_TEMPLATE_ID"))

	fmt.Printf("Response: %+v", responseData)
}
