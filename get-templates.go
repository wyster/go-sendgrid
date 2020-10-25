package main

import (
	"fmt"
	"sendgrid/template"
)

func main() {
	responseData := template.List()
	fmt.Printf("Response: %+v", responseData)
}
