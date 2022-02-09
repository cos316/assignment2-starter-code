# COS316, Assignment 2: HTTP Routing Framework

## Due: 02/23/2022 at 11:00pm

# HTTP Routing Framework

As discussed in lecture, naming schemes are central to system design. 
In this project, you'll build a general library, called an HTTP Routing Framework, 
to help structure web applications based on patterns in end-user requests. 
This general library will support a naming scheme for clients accessing resources 
provided by a web application.

## Getting Started

Before you begin working on this assignment, you will need to set up your
development environment.

Namely, `git clone` your assignment 2 repository (unique to you and your partner) 
from GitHub into whichever local directory you are using to store your
assignment files. (See the *General Assignment Instructions* section toward the bottom
of [this page](https://cos316.princeton.edu/assignments) for more detailed instructions).

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
//           "GET", "PUT", "DELETE". Method strings are case insensitive.
//
// `pattern`: patterns on the request _path_. Patterns can include arbitrary
//           "directories". Directories can include "captures" of the form
//           `:variable_name` that capture the actual directory value into a
//            HTTP query paramter. Leading and trailing '/' (slash) characters
//            are ignored.
//
// Example:
//
//   AddRoute("GET", "/users/:user/recent", RecentUserPosts)
//
// should map all GET requests with a path of the form "/users/*/recent" to the
// `RecentUserPosts` handler. It should populate the query parameter "user" with
// the value of the second directory.
//
// A request of the form "GET /users/cesar/recent HTTP/1.1" will call the
// RecentUserPosts with an `http.Request` with a `URL.RawQuery = "user=cesar"`
//
func (*HttpRouter) AddRoute(method string, pattern string, handler http.HandlerFunc)

// Conforms to the `http.Handler` interface
func (*HttpRouter) ServeHTTP(response http.ResponseWriter, request *http.Request)
```

## Additional Specifications

Be sure that your implementation of the above basic API takes the following into account:

* Your router must support arbitrary paths and HTTP methods, not just those
  required by the microblog client discussed below. You need not (and should not)
  validate that a client-provided method is part of the [official HTTP spec](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods#Specifications).

* You may assume that your router will only be provided with paths that are
  well-formed. That is, paths will be a (possibly empty) list of directory names,
  separated by `/`. Directory names may contain the characters `a-zA-Z0-9_-.`

* HTTP Methods are case-insensitive, whereas paths are case-sensitive. That is,
  `GET` and `get` are equivalent, but `/path/to/file` and `/Path/To/File` are not.

* If (and only if) the request provided to `ServeHTTP` has no associated route,
  your router must write an HTTP "404 Not Found" error as its response. Any response
  with a `Status` field equal to 404 will do.
  You may find Go's [http package](https://golang.org/pkg/net/http/) useful.

* Calling `AddRoute` with a `method`, `pattern` pair that is already associated
  with a handler should update the associated handler to the newly provided value.
  Further, if `pattern` contains a capture, the relevant query parameters
  should be updated. (e.g. `/path/to/:file --> /path/to/:dir`)

* A single directory within a path may contain at most one capture. `/path/to/:file`
  is valid, but `/path/to/:file1:file2` is not.

* A capture must always comprise the entire directory in which it appears. It is
  *not valid* to "prefix" captures, such as in `/path/to/user:id/photos`.

* The empty capture is not valid, either when providing a pattern to AddRoute
  (e.g. `/path/:/file`), or when providing a request path to ServeHTTP
  (e.g. `/path//file`). You may assume your code will never encounter these cases.

* A pattern may include a capture for several values using the same name, as in
  `/path/to/:file/:file`. In this case, the http Response should have a 
  `URL.RawQuery` of `file=<value>&file=<value>`. We impose no restrictions on the 
  order in which key-value pairs appear, except that the values for a particular 
  key must appear in the same order as they did in the request path.
  You may (or may not) find Go's [url package](https://golang.org/pkg/net/url/)
  useful, particularly the [Values type](https://golang.org/pkg/net/url/#Values).

Be aware of the following notable edge case:
* Your router may be asked to route a request for which there are multiple
  applicable routes. Consider a router with a non-capturing route
  `GET /path/to/file` and a capturing route `GET /path/to/:filename`, that has
  been asked to route a request for `GET /path/to/file`.

  In such cases, routes without captures should be given precedence over routes
  with captures. The order in which the routes were added is not considered.

  To be explicit, the route used by the router should be the matching route with
  the largest number of non-capturing path components to the left of its first
  capture. This applies recursively, so that ties are broken by finding the
  matching route with the largest number of non-capturing path components between
  the first and second capture, and so on. If you like, you may informally think
  of this as finding the route which "waits the longest" before capturing each
  of its values.
  
## Performance

A small portion of your grade on this assignment will depend upon your router's 
performance when it contains a large number of routes. For full credit, you should
aim to make your implementation as quick as possible in these cases. 

In particular, if you must iterate over every route in the router to find a matching 
route, your solution will be too slow to earn points for performance. 

Note that the number of points associated with performance is *quite small*, and 
accordingly **you can earn close to full credit** without considering efficiency
at all. 

## Unit Testing

Go provides the [testing package](https://golang.org/pkg/testing/), which is a
convenient framework that allows you to write unit tests for your code to ensure
that it is working correctly and providing the expected results.

For this assignment, you are provided with the file `router_test.go`, which
contains a couple of very simple tests for your router implementation and
demonstrates how to use the `testing` package.

You are encouraged to extend this file to create your own unit tests for your
`http_router` implementation.

You can run your unit tests with the command `go test`, which simply reports the
result of the test, and the reason for failure, if any, or you may add the `-v`
flag to see the verbose output of the unit tests.

For example, run the following from your top-level assignment 2 directory:
```bash
$ go test -v ./http_router
```
Equivalently, you may `cd` into the http_router directory and run the following:
```bash
$ go test -v
```

You will not be graded directly on the quality of your unit tests for this
assignment, but good unit tests will help you debug and understand your
program. We **highly** recommend writing *at least* three or four simple unit tests,
both to familiarize yourself with the API, and to help identify tricky corner
cases or debug unexpected results.

## Sample Application

The assignment starter code includes a sample application that uses the routing
API. Before starting the assignment, your partner and you may want to review this
application source. Understanding this application, and its specific routes, may 
help you design the HTTP Routing Framework library. You may also use this sample 
application as another tool to test your router implementation if you desire, 
but it will not factor into grading in any way.

The application is a simple microblogging application (similar to Twitter).
It uses an in-memory database to store users, threads, and messages, and
presents a JSON-based REST API for listing recent threads from a particular
user or all users the requesting user follows, posting new threads, responding
to threads, and creating new users.

### Running the application

The microblog client and server application can both be compiled using Go's
`go build` command. For example:

```bash
go build -o client ./microblog-client
go build -o server ./microblog-server
```

We have provided a simple Makefile that includes the two commands above.
Feel free to modify the Makefile if you find it convenient.

To build the client and server programs, you can simply run the `make` command,
and the `make` utility will generate two executables, named `client` and `server`.

Once you have built the executables, run each one in a separate terminal window,
in the same way you ran your client and server programs from assignment 1:

```bash
$ ./server
$ ./client recent -user kap
```

If you like, you may also run both programs in the same terminal window by
running the server in the background:

```bash
$ ./server &
$ ./client recent -user kap
```

Note that, unlike the client from assignment 1, this client accepts commands
directly as command-line arguments, rather than by reading from stdin. Run
`./client` with no arguments to view the documentation.

Feel free to read through the source code for the microblogger if you want
more information on how it works, or if you would like to add or alter any
features.

### Using the application

The application uses HTTP basic authentication.

You can exercise the application using a few simple `curl` commands.

## Submission & Grading

Your assignment will be automatically submitted every time you push your changes
to your GitHub repository. Within a couple minutes of your submission, the
autograder will make a comment on your commit listing the output of our testing
suite when run against your code. **Note that you will be graded only on your
changes to the `http_router` package**, and not on your changes to any other files,
though you may modify any files you wish.

You may submit and receive feedback in this way as many times as you like and 
whenever you like.
