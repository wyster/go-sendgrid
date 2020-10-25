package main

import (
	"fmt"
	"sendgrid/template"
)

func main() {
	responseData := template.Create()

	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
