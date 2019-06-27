package main

import (
	"log"
	"net/http"

	"../http_router"
)

type handlerState struct {
	request *http.Request
	writer  http.ResponseWriter
}

type serializedHandler struct {
	c    chan handlerState
	done chan bool
}

func (handler *serializedHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	handler.c <- handlerState{request: request, writer: response}
	<-handler.done
}

func main() {

	DB["amit"] = User{
		Username:  "amit",
		following: []string{"will", "kap"},
		threads:   []Thread{},
	}

	DB["will"] = User{
		Username:  "will",
		following: []string{},
		threads: []Thread{
			Thread{
				ID:        "0",
				Message:   Message{Body: "Hello World", Author: "will"},
				Responses: []Message{},
			},
		},
	}

	DB["kap"] = User{
		Username:  "kap",
		following: []string{},
		threads: []Thread{
			Thread{
				ID:      "1",
				Message: Message{Body: "This is a funny little message", Author: "kap"},
				Responses: []Message{
					Message{
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

	router := http_router.NewRouter()
	router.AddRoute("GET", "/feed", GetFeed)
	router.AddRoute("POST", "/feed", PostToFeed)
	router.AddRoute("POST", "/threads/:thread", FooHandler) // TODO
	router.AddRoute("GET", "/threads/:thread", ThreadResponses)
	router.AddRoute("GET", "/users/:user/recent", GetUserFeed)
	router.AddRoute("POST", "/users/:user/follow", FooHandler) // TODO
	router.AddRoute("POST", "/users/new", NewUser)

	handler := &serializedHandler{}
	handler.c = make(chan handlerState)
	handler.done = make(chan bool)

	go func() {
		for {
			state := <-handler.c
			router.ServeHTTP(state.writer, state.request)
			handler.done <- true
		}
	}()

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error listening for connections: %s", err)
	}
}
