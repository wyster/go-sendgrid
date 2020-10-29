package main

import (
	"fmt"
	"os"
	"sendgrid/template"
)

func main() {
	responseData := template.List(os.Getenv("SENDGRID_TOKEN"))
	fmt.Printf("Response: %+v", responseData)
	for _, id := range responseData.IdsList() {
		fmt.Println(id)
	}
}
