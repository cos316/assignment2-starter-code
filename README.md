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

## Additional Specifications

Be sure that your implementation of the basic API above takes the following into account:
* Your router must support arbitrary paths and HTTP methods, not just those
  required by the microblog client discussed below. You need not (and should not)
  validate that a client-provided method is part of the [official HTTP spec](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods#Specifications).
* HTTP Methods are case-insensitive, whereas paths are case-sensitive. That is,
  `GET` and `get` are equivalent, but `/path/to/file` and `Path/To/File` are not.
* Paths should be indifferent to leading and trailing slashes. `/path/` and `path`
  are equivalent.
* If (and only if) the request provided to `ServeHTTP` has no associated route,
  your router must write an HTTP "404 Not Found" error to the response. Any response
  with a `Status` field equal to 404 will do.
  You may find Go's [http package](https://golang.org/pkg/net/http/) useful.
* Calling `AddRoute` with a `method`, `pattern` pair that are already associated
  with a handler should update the associated handler to the newly provided value.
  Further, if the `pattern` contains a capture, the relevant query parameters
  should be updated. (e.g. `/path/to/:file --> /path/to/:dir`)
* A single directory within a path may contain at most one capture. `/path/to/:file`
  is valid, but `/path/to/:file1:file2` is not.
* A capture must always comprise the entire directory in which it appears. It is
  *not valid* to "prefix" captures, such as in `/path/to/user:id/photos`.

Be aware of the following notable edge case:
* It may be the case that there are several routes which could apply to a given
  path. For example, it is valid to add a non-capturing route `GET /path/to/file`
  and a capturing route `GET /path/to/:filename`. In such cases, non-capturing routes
  should be given precedence over capturing ones for the purposes of resolving matches.
  In general, favor the resolution which is most specific and which shadows (makes
  inaccessible) the smallest set of paths.
  ***TODO: There are situations with more nuanced ambiguities, but we don't have
  tests for these (yet - coming soon!), so ignore them for now.***

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
