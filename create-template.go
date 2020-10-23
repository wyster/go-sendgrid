package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var SendgridApiToken = os.Getenv("SENDGRID_TOKEN")
	fmt.Println(SendgridApiToken)

	req, err := http.NewRequest(
		"POST",
		"https://api.sendgrid.com/v3/templates",
		strings.NewReader(`
			{
				"name":"test",
				"generation":"dynamic"
			}
		`),
	)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+SendgridApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	var responseData struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Generation string `json:"generation"`
		UpdatedAt  string `json:"updated_at"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
