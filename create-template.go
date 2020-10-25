package main

import (
	"fmt"
	"os"
	"sendgrid/template"
)

func main() {
	responseData := template.Create(os.Getenv("SENDGRID_TOKEN"))

	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
