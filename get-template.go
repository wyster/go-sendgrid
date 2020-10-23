package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var SendgridApiToken = os.Getenv("SENDGRID_TOKEN")
	var SendgridTemplateId = os.Getenv("SENDGRID_TEMPLATE_ID")
	fmt.Println(SendgridApiToken)

	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates/"+SendgridTemplateId,
		strings.NewReader(""),
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
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Response status:", resp.Status)
	//fmt.Println(string(bodyBytes))

	type Versions []struct {
		Id         string `json:"id"`
		UserId     int    `json:"user_id"`
		TemplateId string `json:"template_id"`
		Active     int    `json:"active"`
	}

	var responseData struct {
		Id         string   `json:"id"`
		Name       string   `json:"name"`
		Generation string   `json:"generation"`
		UpdatedAt  string   `json:"updated_at"`
		Versions   Versions `json:"versions"`
	}

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Println(responseData.Id)
	fmt.Printf("Response: %+v", responseData)
}
