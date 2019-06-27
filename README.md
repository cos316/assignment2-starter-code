# HTTP Routing Framework

In this project, you'll build a library to help structure web applications based
on patterns in end-user requests.

## API

Your solution must implement the following API:

```go
type HttpRouter struct {
    // this can include whatever fields you want
}

// Creates a new HttpRouter
func NewRouter() *HttpRouter

// Adds a new route to the HttpRouter
//
// Routes any request that matches the given `method` and `pattern` to a
// `handler`.
//
// `method`: should support arbitrary method strings, and least each of "POST",
//           "GET", "PUT", "DELETE". Method strings are case insensitive
//
// `pattern`: patterns on the request _path_. Patterns can include arbitrary
//           "directories". Directories can include "captures" of the form
//           `:variable_name` that capture the actual directory value into a
//            HTTP query paramter. Leading and trailing '/' (slash) characters
//            are ignored
//
// Example:
//
// AddRoute("GET", "/users/:user/recent", RecentUserPosts)
//
// Should map all GET requests with a path of the form "/users/*/recent" to the
// `RecentUserPosts` handler. It should populate the query parameter "user" with
// the value of the second directory.
//
// A request of the form "GET /users/cesar/recent HTTP/1.1" will call the
// RecentUserPosts with an `http.Request` with a `RawQuery = "user=cesar"`
//
func (*HttpRouter) AddRoute(method string, pattern string, handler http.HandlerFunc)

// Conforms to the `http.HandlerHttp` interface
func (*HttpRouter) ServeHTTP(response http.ResponseWriter, request *http.Request)
```

## Sample Application

The assignment starter code includes a sample application that uses the routing
API. The application is a simple microblogging application (similar to
Twitter).

It uses an in-memory database to store users, threads, and messages, and
presents a JSON-based REST API for listing recents threads from a particular
user or all users the requesting user follows, posting new threads, responding
to threads, and creating new users.

### Running the application

Compile the application Go's `build` subcommand:

```bash
$ go build
$ ./project2
```

Or run it directly with Go's `run` subcommand:

```bash
$ go run
Listening for http on port :8080
http://localhost:8080
```

### Using the application

The application uses HTTP basic authentication

You can excercise the application using a few simple `curl` commands.



