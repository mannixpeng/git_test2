package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", errors.Wrap(err, "unable to complete Get request")
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "unable to read response data")
	}

	return string(out), nil
}

func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()
	t.Log("svr.URL: ", svr.URL)
	c := NewClient(svr.URL)
	res, err := c.UpperCase("anything")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	// res: expected\r\n
	// due to the http protocol cleanup response
	res = strings.TrimSpace(res)
	if res != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
