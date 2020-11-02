package main

import (
	"fmt"
	"github.com/wyster/go-sendgrid/template"
	"log"
	"os"
)

func main() {
	responseData, err := template.List(os.Getenv("SENDGRID_TOKEN"))
	if err != nil {
		log.Fatal("Templates not fetched! ", err)
	}
	fmt.Printf("Response: %+v", responseData)
	for _, id := range responseData.IdsList() {
		fmt.Println(id)
	}
}
