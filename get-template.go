package main

import (
	"fmt"
	"sendgrid/template"
)

func main() {
	responseData := template.Get()

	fmt.Printf("Response: %+v", responseData)
}
