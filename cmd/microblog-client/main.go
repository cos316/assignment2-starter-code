package main

import (
	"flag"
)

func main() {
	app := &app{commands: make(map[string]command)}
	app.commands["feed"] =
		newCommand("Fetch the requesting user's recent feed", feed)
	app.commands["new"] =
		newCommand("Create a new user", newUser)
	app.commands["recent"] =
		newCommand("List the user's recent posts", userFeed)
	app.commands["follow"] =
		newCommand("Follow a user from the requesting user", follow)
	app.commands["post"] =
		newCommand("Post a new thread", postThread)
	app.commands["help"] =
		newCommand("Displays this help message", help)

	flag.StringVar(&app.username, "user", "", "The requesting username, e.g. \"cesar\"")
	flag.Parse()

	app.run(flag.Arg(0))
}
