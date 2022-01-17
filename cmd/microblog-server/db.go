package main

var GlobalCounter int64 = 11

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

type Thread struct {
	ID        string    `json:"id"`
	Message   Message   `json:"message"`
	Responses []*Message `json:"responses"`
}

type User struct {
	Username  string `json:"username"`
	following []string
	threads   []*Thread
}

var DB map[string]*User = make(map[string]*User)
