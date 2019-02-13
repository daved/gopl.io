package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Err ...
var (
	ErrRepoNotSet = fmt.Errorf("owner and repo must be set to use this function")
	ErrNoBody     = fmt.Errorf("request requires body; likely library error")
)

func sendDecode(method, url string, hs headers, data, v interface{}) error {
	if missingBody(method, data) {
		return ErrNoBody
	}

	var body io.Reader
	if data != nil {
		bs, err := json.Marshal(&data)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(bs)
	}

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	hs.apply(r.Header)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint

	if !statusSuccess(res.StatusCode) {
		return fmt.Errorf("query failed: %s", res.Status)
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(&v)
}

func missingBody(method string, data interface{}) bool {
	return data == nil && method != http.MethodGet && method != http.MethodDelete
}

func statusSuccess(code int) bool {
	return code > 199 && code < 300
}
