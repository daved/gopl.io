package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

func unmarshalStored(filename string, v interface{}) error {
	efmt := "unmarshal stored: %s"

	f, err := os.Open(path.Clean(filename))
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf(efmt, err)
		}

		return nil
	}
	defer f.Close() //nolint

	st, err := f.Stat()
	if err != nil {
		return fmt.Errorf(efmt, err)
	}

	if st.Size() == 0 {
		return nil
	}

	if err = json.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf(efmt, err)
	}

	return nil
}

func marshalStored(filename string, v interface{}) error {
	efmt := "marshal stored: %s"

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0660) //nolint
	if err != nil {
		return fmt.Errorf(efmt, err)
	}
	defer f.Close() //nolint

	if err = json.NewEncoder(f).Encode(v); err != nil {
		return fmt.Errorf(efmt, err)
	}

	if err = f.Sync(); err != nil {
		return fmt.Errorf(efmt, err)
	}

	return nil
}

func sendDecode(url string, v interface{}) error {
	efmt := "send decode: %s"

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}
	defer res.Body.Close() //nolint

	if !statusSuccess(res.StatusCode) {
		return fmt.Errorf(efmt, "query: "+res.Status)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return fmt.Errorf(efmt, err)
	}

	return nil
}

func statusSuccess(code int) bool {
	return code > 199 && code < 300
}

func isExpired(expiry time.Duration, target, test time.Time) bool {
	return test.Sub(target) > expiry
}
