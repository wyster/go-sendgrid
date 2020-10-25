package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
)

func main() {
	responseData := template.Create()

	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
