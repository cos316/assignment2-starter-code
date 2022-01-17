// microblog-server on localhost:8080
//
// Useage: microblog-server

package main

import (
	"log"
	"net/http"

	"cos316.princeton.edu/assignment2/http_router"
)

// define struct to store the HTTP request and header
type handlerState struct {
	request *http.Request
	writer  http.ResponseWriter
}

// we want to serialize requests, so define a struct to
// for sending handlerState struct
type serializedHandler struct {
	c    chan handlerState
	done chan bool
}

// define a ServeHTTP method for the serializeHandler
// we can use this to define a customer handler for ListenAndServe
func (handler *serializedHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// ListenAndServe will call this function when a client makes an HTTP request
	handler.c <- handlerState{request: request, writer: response} // send request and writer to handler channel
	<-handler.done                                                // block until done
}

func main() {
	// initlaize database (DB) with some users and users' data
	DB["amit"] = &User{
		Username:  "amit",
		following: []string{"will", "kap"},
		threads:   []*Thread{},
	}

	DB["will"] = &User{
		Username:  "will",
		following: []string{},
		threads: []*Thread{
			&Thread{
				ID:        "0",
				Message:   Message{Body: "Hello World", Author: "will"},
				Responses: []*Message{},
			},
		},
	}

	DB["kap"] = &User{
		Username:  "kap",
		following: []string{},
		threads: []*Thread{
			&Thread{
				ID:      "1",
				Message: Message{Body: "This is a funny little message", Author: "kap"},
				Responses: []*Message{
					&Message{
						Body:   "Interesting...",
						Author: "amit",
					},
				},
			},
		},
	}

	GlobalCounter = 2
	/*
		Microblogging service

		GET /feed -> recent
		POST /feed -> post new message
		POST /feed/:thread -> post response to message
		GET /threads/:thread -> responses to thread
		GET /users/:user
		GET /users/:user/recent
		POST /users/:user/follow
		POST /users/new
	*/

	// create a new router (from the general libary)
	// add the routes (names) we want to support
	router := http_router.NewRouter()
	router.AddRoute("GET", "/feed", GetFeed)
	router.AddRoute("POST", "/feed", PostToFeed)
	router.AddRoute("POST", "/threads/:thread", RespondToThread)
	router.AddRoute("GET", "/threads/:thread", ThreadResponses)
	router.AddRoute("GET", "/users/:user/recent", GetUserFeed)
	router.AddRoute("POST", "/users/:user/follow", FollowUser)
	router.AddRoute("POST", "/users/new", NewUser)

	// set up the custom handler
	handler := &serializedHandler{}
	handler.c = make(chan handlerState)
	handler.done = make(chan bool)

	// Use a go routine wait to serve HTTP requests
	go func() {
		for {
			state := <-handler.c                          // send data to state channel
			router.ServeHTTP(state.writer, state.request) // handle the route
			handler.done <- true                          // signal done to channel
		}
	}()

	// listening on localhost:8080
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error listening for connections: %s", err)
	}
}
