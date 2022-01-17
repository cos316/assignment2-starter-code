package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func newUser(app *app) {
	response, err := postForm(app, "users/new", url.Values{"username": {app.username}}, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response.Status)
	defer response.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(response.Body).Decode(body)
	fmt.Println(body)
}

func follow(app *app) {
	console := bufio.NewScanner(os.Stdin)
	fmt.Printf("Who do you want to follow? ")
	console.Scan()
	user := console.Text()

	response, err := postForm(app, fmt.Sprintf("users/%s/follow", user), nil, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response.Status)
	defer response.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(response.Body).Decode(body)
	fmt.Println(body)
}

func postThread(app *app) {
	console := bufio.NewScanner(os.Stdin)
	fmt.Printf("What message do you want to post? ")
	console.Scan()
	message := console.Text()

	response, err := postForm(app, "feed", url.Values{"body": {message}}, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response.Status)
	defer response.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(response.Body).Decode(body)
	fmt.Println(body)
}

func userFeed(app *app) {
	url := fmt.Sprintf("http://localhost:8080/users/%s/recent", app.username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response.Status)
	defer response.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(response.Body).Decode(body)
	fmt.Println(body)
}

func feed(app *app) {
	request, err := http.NewRequest("GET", "http://localhost:8080/feed", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	request.SetBasicAuth(app.username, "nopassword")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response.Status)
	fmt.Printf("==============================\n\n")
	var body []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, thread := range body {
		msg := thread["message"].(map[string]interface{})
		fmt.Printf("(%s) @%s: \n\t%s\n\n", thread["id"], msg["author"], msg["body"])
	}
}
