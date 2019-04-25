package kutt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// URLDefinition represent URL definition from API
type URLDefinition struct {
	ID       string `json:"id"`
	Target   string `json:"target"`
	ShortURL string `json:"shortUrl"`
}

// ResponseErr will handling error response from API
type ResponseErr struct {
	Result string `json:"error"`
}

func (e ResponseErr) Error() string {
	var response string
	if strings.Contains(e.Result, "Unauthorized") {
		response = fmt.Sprintf("Error: You are %s Consider to check your API TOKEN", e.Result)
	} else {
		response = fmt.Sprintf("Error: %s", e.Result)
	}
	return response
}

// API represent KuttIt API service
type API struct {
	httpClient   *http.Client
	Timeout      time.Duration
	BaseURL      string
	Token        string
	CustomDomain string
}

// Call represent process making HTTP call on API
func (api *API) Call(method, path string, payload interface{}) (*http.Response, error) {
	u, _ := url.Parse(api.BaseURL)
	u.Path = path

	var buf []byte
	if payload != nil {
		json, _ := json.Marshal(payload)
		buf = json
	}

	req, err := http.NewRequest(
		method,
		u.String(),
		bytes.NewBuffer(buf),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", api.Token)

	api.httpClient = &http.Client{Timeout: api.Timeout}
	return api.httpClient.Do(req)
}

// GetListURL will return list of URL from API
func (api *API) GetListURL() ([]URLDefinition, error) {
	response, err := api.Call(http.MethodGet, "/api/url/geturls", nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		var responseErr ResponseErr
		json.NewDecoder(response.Body).Decode(&responseErr)
		return []URLDefinition{}, responseErr
	}

	var urlResult struct {
		List     []URLDefinition `json:"list"`
		CountAll int64           `json:"countAll"`
	}

	err = json.NewDecoder(response.Body).Decode(&urlResult)
	if err != nil {
		return nil, err
	}
	return urlResult.List, nil
}

// SubmitURL will submit long url and return short url
// with some optional customization if needed
func (api *API) SubmitURL(longURL, customURL, password string, reuse bool) (URLDefinition, error) {
	payload := struct {
		Target    string `json:"target"`
		CustomURL string `json:"customurl"`
		Password  string `json:"password"`
		Reuse     bool   `json:"reuse"`
	}{}

	payload.Target = longURL

	if customURL != "" {
		payload.CustomURL = customURL
	}

	if password != "" {
		payload.Password = password
	}

	if !reuse {
		payload.Reuse = reuse
	}

	response, err := api.Call(http.MethodPost, "/api/url/submit", payload)
	if err != nil {
		return URLDefinition{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		var responseErr ResponseErr
		json.NewDecoder(response.Body).Decode(&responseErr)
		return URLDefinition{}, responseErr
	}

	var url URLDefinition
	err = json.NewDecoder(response.Body).Decode(&url)
	if err != nil {
		return URLDefinition{}, err
	}
	return url, nil
}

// DeleteURL will return nothing if url successfully deleted
func (api *API) DeleteURL(targetURL string) error {
	var id string
	if strings.ContainsAny(targetURL, "/") {
		splitTarget := strings.Split(targetURL, "/")
		id = splitTarget[len(splitTarget)-1]
	} else {
		id = targetURL
	}

	payload := struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	}{}

	payload.ID = id

	if api.CustomDomain != "" {
		payload.Domain = api.CustomDomain
	}

	response, err := api.Call(http.MethodPost, "/api/url/deleteurl", payload)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		var responseErr ResponseErr
		json.NewDecoder(response.Body).Decode(&responseErr)
		return responseErr
	}
	return nil
}
