package hclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Do executes an HTTP request and returns the response body as a string.
// Non-200 response codes will be returned as an error with the response body.
func DoString(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	return string(body), nil
}

// call HTTP request with raw json
func Do(req *http.Request) (result json.RawMessage, err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	result = body
	return result, nil
}
