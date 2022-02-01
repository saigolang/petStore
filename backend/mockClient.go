package backend

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type MockClient struct {
	mockResponse string
	statusCode   int
	httpError    error
}

func (m MockClient) Do(req *http.Request) (*http.Response, error) {

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(m.mockResponse)))
	return &http.Response{
		StatusCode: m.statusCode,
		Body:       responseBody,
	}, m.httpError
}
