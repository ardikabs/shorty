package kutt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	expectedMethod := http.MethodPost

	type payload struct {
		CustomURL string `json:"customurl"`
		Target    string `json:"target"`
	}

	expectedBody := payload{
		CustomURL: "custom-url",
		Target:    "https://target.url",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Errorf("expected with method %s, but got method %s", expectedMethod, r.Method)
		}

		reqbody, _ := ioutil.ReadAll(r.Body)
		var body payload
		json.Unmarshal(reqbody, &body)

		if !reflect.DeepEqual(body, expectedBody) {
			t.Errorf("expected %v but got %v", expectedBody, body)
		}

		w.Write([]byte(`{
				"count": 2,
				"createdAt": "2019-04-22T07:52:59.973Z",
				"id": "learning",
				"target": "https://www.packtpub.com/packt/offers/free-learning",
				"password": false,
				"shortUrl": "http://urls.ardikabs.id/learning"
			}
		}`))
	}
	s := httptest.NewServer(http.HandlerFunc(handler))

	defer s.Close()

	api := API{
		BaseURL: s.URL,
		Token:   "true-api-token",
		Timeout: 5 * time.Second,
	}

	customPayload := payload{
		CustomURL: "custom-url",
		Target:    "https://target.url",
	}

	api.Call(http.MethodPost, "/api/test", customPayload)
}

func TestGetListURL(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"list": [
				{
					"count": 2,
					"createdAt": "2019-04-22T07:52:59.973Z",
					"id": "learning",
					"target": "https://www.packtpub.com/packt/offers/free-learning",
					"password": false,
					"shortUrl": "http://urls.ardikabs.id/learning"
				},
				{
					"count": 3,
					"createdAt": "2019-04-22T07:48:58.696Z",
					"id": "bukalapak",
					"target": "https://bukalapak.com",
					"password": false,
					"shortUrl": "http://urls.ardikabs.id/bukalapak"
				}
			],
			"countAll": 2
		}`))
	}
	s := httptest.NewServer(http.HandlerFunc(handler))

	defer s.Close()

	api := API{
		BaseURL: s.URL,
		Token:   "true-api-token",
	}

	urls, err := api.GetListURL()
	want := 2

	if err != nil {
		t.Errorf("expected to not get any error, but got %s", err)
	}

	if len(urls) != want {
		t.Errorf("expected %d data, got %d data", want, len(urls))
	}

}

func TestSubmitURL(t *testing.T) {
	t.Run("custom url with success response", func(t *testing.T) {
		s := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"count": 0,
					"createdAt": "2019-04-22T07:52:59.973Z",
					"id": "my-google",
					"target": "https://google.com",
					"password": false,
					"shortUrl": "http://example.com/my-google"
				}`))
			}),
		)

		defer s.Close()

		api := API{
			BaseURL: s.URL,
			Token:   "true-api-token",
		}

		customURL := "my-google"
		got, err := api.SubmitURL(
			"https://google.com",
			customURL,
			"",
			false,
		)

		want := fmt.Sprintf("http://example.com/%s", customURL)
		if err != nil {
			t.Errorf("expected to not get any error, but got %s", err)
		}

		if got.ShortURL != want {
			t.Errorf("expected %s, got %s", want, got.ShortURL)
		}
	})

	t.Run("custom url with unsuccessful response", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{
				"error": "Error from server"
			}`))
		}

		s := httptest.NewServer(http.HandlerFunc(handler))
		defer s.Close()

		api := API{
			BaseURL: s.URL,
			Token:   "true-api-token",
		}

		customURL := "my-google"
		_, err := api.SubmitURL(
			"watt://uglyface.com",
			customURL,
			"",
			false,
		)

		if err == nil {
			t.Errorf("expected to get error, but got nothing")
		}
	})

}

func TestDeleteURL(t *testing.T) {
	t.Run("successfull delete", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		defer s.Close()

		api := API{
			BaseURL: s.URL,
			Token:   "true-api-token",
		}

		err := api.DeleteURL("https://kutt.it/my-google")

		if err != nil {
			t.Errorf("expected to not get any error, but got %s", err)
		}
	})

	t.Run("unsuccessfull delete", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		defer s.Close()

		api := API{
			BaseURL: s.URL,
			Token:   "true-api-token",
		}

		err := api.DeleteURL("https://kutt.it/my-google")

		if err == nil {
			t.Errorf("expected to get error, but got nothing")
		}
	})

	t.Run("custom domain provided but unsuccessful delete url", func(t *testing.T) {
		customDomain := "custom.example.com"

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				ID           string `json:"id"`
				CustomDomain string `json:"domain"`
			}{}
			json.NewDecoder(r.Body).Decode(&payload)

			if payload.CustomDomain != customDomain {
				t.Errorf("expected %s but got %s", customDomain, customDomain)
			}

		}
		s := httptest.NewServer(http.HandlerFunc(handler))
		defer s.Close()

		api := API{
			BaseURL:      s.URL,
			Token:        "true-api-token",
			CustomDomain: customDomain,
		}

		err := api.DeleteURL("https://kutt.it/my-google")

		if err == nil {
			t.Errorf("expected to get error, but got nothing")
		}
	})

}
