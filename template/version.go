package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetVersion(ApiToken string, TemplateId string, TemplateVersion string) (Version, error) {
	var responseData Version

	req, err := http.NewRequest(
		"GET",
		"https://api.sendgrid.com/v3/templates/"+TemplateId+"/versions/"+TemplateVersion,
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

func CreateVersion(ApiToken string, TemplateId string, Name string) (Version, error) {
	var responseData Version

	req, err := http.NewRequest(
		"POST",
		"https://api.sendgrid.com/v3/templates/"+TemplateId+"/versions",
		strings.NewReader(fmt.Sprintf(`
			{
				"template_id": "%s",
				"active": 1,
				"name": "%s"
			}
		`, TemplateId, Name)),
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
