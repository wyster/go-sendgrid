package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
	"log"
	"os"
)

func main() {
	responseData, err := template.Create(os.Getenv("SENDGRID_TOKEN"))
	if err != nil {
		log.Fatal("Template not created! ", err)
	}
	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
