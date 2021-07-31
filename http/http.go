package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTP struct {
	GetURL string
}

type HTTPer interface {
	Get(path string) ([]byte, error)
}

func NewHTTP(getURL string) HTTPer {
	return &HTTP{
		GetURL: getURL,
	}
}

func (h *HTTP) Get(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", h.GetURL, path))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
