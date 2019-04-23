package kutt

import (
	"bytes"
	"encoding/json"
	"errors"
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

// API represent KuttIt API service
type API struct {
	httpClient   *http.Client
	BaseURL      *url.URL
	APIToken     string
	CustomDomain string
}

// GetListURL will return list of URL from API
func (api *API) GetListURL() ([]URLDefinition, error) {

	rel := &url.URL{Path: "/api/url/geturls"}
	u := api.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Key", api.APIToken)

	api.httpClient = &http.Client{Timeout: 15 * time.Second}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var urlResult struct {
		List     []URLDefinition `json:"list"`
		CountAll int64           `json:"countAll"`
	}

	err = json.NewDecoder(resp.Body).Decode(&urlResult)
	if err != nil {
		return nil, err
	}
	return urlResult.List, nil
}

// SubmitURL will submit long url and return short url
// with some optional customization if needed
func (api *API) SubmitURL(longURL, customURL, password string, reuse bool) (URLDefinition, error) {
	rel := &url.URL{Path: "/api/url/submit"}

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

	body, err := json.Marshal(payload)
	if err != nil {
		return URLDefinition{}, err
	}

	u := api.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return URLDefinition{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", api.APIToken)

	api.httpClient = &http.Client{Timeout: 15 * time.Second}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return URLDefinition{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return URLDefinition{}, errors.New(resp.Status)
	}

	var url URLDefinition
	err = json.NewDecoder(resp.Body).Decode(&url)
	if err != nil {
		return URLDefinition{}, err
	}
	return url, nil
}

// DeleteURL will return nothing if url successfully deleted
func (api *API) DeleteURL(targetURL string) error {
	rel := &url.URL{Path: "/api/url/deleteurl"}

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

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	u := api.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", api.APIToken)

	api.httpClient = &http.Client{Timeout: 15 * time.Second}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
