package github

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type header struct {
	key, val string
}

func newHeaderAuth(userURL, username, token string) (*header, error) {
	creds := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))

	h := header{
		"Authorization",
		"Basic " + creds,
	}

	r, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set(h.key, h.val)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //nolint

	if !statusSuccess(res.StatusCode) {
		return nil, fmt.Errorf("cannot verify username and token: %s", res.Status)
	}

	return &h, nil
}

type headers []*header

func (hs headers) apply(hh http.Header) {
	for _, h := range hs {
		hh.Set(h.key, h.val)
	}
}
