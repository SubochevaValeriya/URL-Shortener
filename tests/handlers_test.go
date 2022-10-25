package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddURLHandler(t *testing.T) {
	t.Parallel()
	type testCase struct {
		url  string
		code int
	}

	testCases := []testCase{
		{"https://github.com/SubochevaValeriya", 303},
		{"829038jdksjm,x", 400},
	}

	handler, _ := initForTest()

	for _, testCase := range testCases {

		request, err := http.NewRequest("POST", "/new?longURL="+testCase.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()

		handler.AddURL(responseRecorder, request)
		status := responseRecorder.Code

		if status != testCase.code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.code)
		}
	}
}

func TestGetURLHandler(t *testing.T) {
	t.Parallel()
	handler, repo := initForTest()

	addURLForTest(repo)

	type testCase struct {
		url  string
		code int
	}

	testCases := []testCase{
		{"dhfh3", 303},
		{"", 400},
		{"notInTheBase", 404},
	}

	for _, testCase := range testCases {
		request, err := http.NewRequest("GET", "/urls/"+testCase.url, nil)

		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()

		handler.GetURL(responseRecorder, request)

		status := responseRecorder.Code

		if status != testCase.code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.code)
		}
	}

	deleteURLForTest(repo)
}
