/******************************************************************************
 *  router_test.go
 *  Author:
 *  Usage:    `go test`  or  `go test -v`
 *  Description:
 *    An incomplete unit testing suite for router.go. You are welcome to change
 *    anything in this file however you would like. You are strongly encouraged
 *    to create additional tests for your implementation, as the ones provided
 *    here are extremely basic, and intended only to demonstrate how to test
 *    your program.
 ******************************************************************************/

package http_router

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/******************************************************************************/
/*                                Constants                                   */
/******************************************************************************/

// An incomplete list of possible HTTP methods. See http docs for more.
const (
	httpGet  = http.MethodGet
	httpPost = http.MethodPost
	httpPut  = http.MethodPut
)

// An incomplete list of possible HTTP status codes. See http docs for more.
const (
	httpOK       = http.StatusOK       // 200
	httpNotFound = http.StatusNotFound // 404
)

/******************************************************************************/
/*                                 Helpers                                    */
/******************************************************************************/

// responseBodyToString consumes the body of response, returning it as a string
func responseBodyToString(response *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}

// You are encouraged to add helper functions here to make testing easier.

/******************************************************************************/
/*                                 Handlers                                   */
/******************************************************************************/

// echoPathHandler writes the path of the request to the response body, but
// otherwise ignores all other request fields.
func echoPathHandler(response http.ResponseWriter, request *http.Request) {
	pathBytes := []byte(request.URL.Path)
	response.Write(pathBytes)
}

// echoMethodHandler writes the method of the request to the response body, but
// otherwise ignores all other request fields.
func echoMethodHandler(response http.ResponseWriter, request *http.Request) {
	methodBytes := []byte(request.Method)
	response.Write(methodBytes)
}

/******************************************************************************/
/*                                  Tests                                     */
/******************************************************************************/

// TestBasicGet creates a new router, registers a single route with it,
// and then creates and sends a single request to the router. After the router
// returns a response, the fields of the response are inspected.
func TestBasicGet(t *testing.T) {
	router := NewRouter()
	router.AddRoute(httpGet, "/index.html", echoMethodHandler)

	request := httptest.NewRequest(httpGet, "http://localhost:8080/index.html", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	if response.StatusCode != httpOK {
		t.Errorf("Test failed: Router gave non-200 status code: %d", response.StatusCode)
	}

	body := responseBodyToString(response)
	if body != httpGet {
		t.Errorf("Test failed: Expected %s and received %s.", httpGet, body)
	}
}

// TestBasicPaths creates a new router, registers several routes with it,
// and then creates and sends a single request to the router for each route.
// After the router returns a response, the fields of the response are inspected.
func TestBasicPaths(t *testing.T) {
	paths := []string{
		"/index.html",
		"/students/index.html",
		"/assignments/index.html",
		"/syllabus.html",
	}

	router := NewRouter()

	for _, path := range paths {
		router.AddRoute(httpGet, path, echoPathHandler)
	}

	for _, path := range paths {
		request := httptest.NewRequest(httpGet, "http://localhost:8080" + path, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		if response.StatusCode != httpOK {
			t.Errorf("Test failed: Router gave non-200 status code: %d", response.StatusCode)
		}

		body := responseBodyToString(response)
		if body != path {
			t.Errorf("Test failed: Expected %s and received %s.", path, body)
		}
	}
}
