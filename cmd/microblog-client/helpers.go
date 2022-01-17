package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func postForm(app *app, path string, values url.Values, auth bool) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8080/%s", path)
	var body io.Reader
	if values == nil {
		body = nil
	} else {
		body = strings.NewReader(values.Encode())
	}
	request, err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if auth {
		request.SetBasicAuth(app.username, "nopassword")
	}
	return client.Do(request)
}
