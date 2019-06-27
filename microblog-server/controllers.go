package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type constrainedThread struct {
	ID      string  `json:"id"`
	Message Message `json:"message"`
}

func GetFeed(response http.ResponseWriter, request *http.Request) {
	callingUser, _, ok := request.BasicAuth()
	if !ok {
		callingUser = "temp"
	}

	user, ok := DB[callingUser]
	result := make([]constrainedThread, 0)
	if ok {
		for _, following := range user.following {
			friend, ok := DB[following]
			if ok {
				threads := make([]constrainedThread, len(friend.threads))
				for i, t := range friend.threads {
					threads[i] = constrainedThread{
						ID:      t.ID,
						Message: t.Message,
					}
				}
				result = append(result, threads...)
			}
		}
	} else {
		http.Error(response, "Not authorized", http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func GetUserFeed(response http.ResponseWriter, request *http.Request) {
	callingUser := request.URL.Query().Get("user")

	user, ok := DB[callingUser]
	if ok {
		result := make([]constrainedThread, len(user.threads))
		for i, t := range user.threads {
			result[i] = constrainedThread{
				ID:      t.ID,
				Message: t.Message,
			}
		}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	} else {
		http.Error(response, "Not authorized", http.StatusInternalServerError)
		return
	}
}

func PostToFeed(response http.ResponseWriter, request *http.Request) {
	callingUser, _, ok := request.BasicAuth()
	if !ok {
		callingUser = "temp"
	}

	user, ok := DB[callingUser]
	if ok {
		request.ParseForm()
		values := request.Form
		threadID := strconv.FormatInt(GlobalCounter, 10)
		GlobalCounter++
		thread := Thread{
			ID: threadID,
			Message: Message{
				Body:   values.Get("body"),
				Author: callingUser,
			},
			Responses: []Message{},
		}
		user.threads = append([]Thread{thread}, user.threads...)
		DB[callingUser] = user
		js, _ := json.Marshal(thread)
		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	} else {
		http.Error(response, "Not authorized", http.StatusInternalServerError)
		return
	}

}

func ThreadResponses(response http.ResponseWriter, request *http.Request) {
	threadID := request.URL.Query().Get("thread")
	for _, user := range DB {
		for _, thread := range user.threads {
			if thread.ID == threadID {
				js, _ := json.Marshal(thread)
				response.Header().Set("Content-Type", "application/json")
				response.Write(js)
				return
			}
		}
	}
}

func NewUser(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	values := request.Form
	username := values.Get("username")
	if _, ok := DB[username]; !ok {
		user := User{
			Username:  username,
			threads:   []Thread{},
			following: []string{},
		}
		DB[username] = user
		js, _ := json.Marshal(user)
		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	}
}

func FooHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Bar")
}
