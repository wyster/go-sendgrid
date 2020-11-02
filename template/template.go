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

func Create(ApiToken string) (Template, error) {
	var responseData Template

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
		return responseData, &Error{0, "Error reading request. " + err.Error()}
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return responseData, &Error{0, "Error reading response. " + err.Error()}
	}
	if resp.StatusCode != 201 {
		return responseData, &Error{resp.StatusCode, resp.Status}
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		return responseData, &Error{resp.StatusCode, "Error reading response. " + err.Error()}
	}

	return responseData, nil
}

func Get(ApiToken string, TemplateId string) (Template, error) {
	var responseData Template

	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates/"+TemplateId,
		strings.NewReader(""),
	)
	if err != nil {
		return responseData, &Error{0, "Error reading request. " + err.Error()}
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return responseData, &Error{0, "Error reading response. " + err.Error()}
	}
	if resp.StatusCode != 200 {
		return responseData, &Error{resp.StatusCode, resp.Status}
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		return responseData, &Error{resp.StatusCode, "Error reading response. " + err.Error()}
	}

	return responseData, nil
}

func List(ApiToken string) (Templates, error) {
	var responseData Templates

	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates?generations=dynamic",
		strings.NewReader(""),
	)
	if err != nil {
		return responseData, &Error{0, "Error reading request. " + err.Error()}
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+ApiToken)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return responseData, &Error{0, "Error reading response. " + err.Error()}
	}
	if resp.StatusCode != 200 {
		return responseData, &Error{resp.StatusCode, resp.Status}
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		return responseData, &Error{resp.StatusCode, "Error reading response. " + err.Error()}
	}

	return responseData, nil
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
