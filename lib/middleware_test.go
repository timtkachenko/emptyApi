package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestLogging(t *testing.T) {
	a := assert.New(t)

	ts := httptest.NewServer(RequestLogging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	a.Equal("Hello, client\n", string(greeting))
}
