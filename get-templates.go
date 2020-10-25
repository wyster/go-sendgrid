package main

import (
	"fmt"
	"os"
	"sendgrid/template"
)

func main() {
	responseData := template.List(os.Getenv("SENDGRID_TOKEN"))
	fmt.Printf("Response: %+v", responseData)
}
