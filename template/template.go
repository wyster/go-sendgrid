package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Create(ApiToken string) Template {
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
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var responseData Template

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	return responseData
}

func Get(ApiToken string, TemplateId string) Template {
	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates/"+TemplateId,
		strings.NewReader(""),
	)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var responseData Template

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	return responseData
}

func List(ApiToken string) Templates {
	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates?generations=dynamic",
		strings.NewReader(""),
	)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var responseData Templates

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	return responseData
}

func Delete(ApiToken string, TemplateId string) (bool, error) {
	req, err := http.NewRequest(
		"DELETE",
		"https://api.sendgrid.com/v3/templates/"+TemplateId,
		strings.NewReader(""),
	)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return true, nil
	}

	return false, &Error{resp.StatusCode, resp.Status}
}
