package example

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	//create reqeust
	req, err := http.NewRequest(http.MethodGet, "localhost:8888/double?v=3", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	// create recorder as response writer
	rec := httptest.NewRecorder()
	doubleHandler(rec, req)
	// extract response from recorder
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected response status %d; but got %d", http.StatusOK, res.StatusCode)
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not get response body: %v", err)
	}

	if num, err := strconv.Atoi(strings.TrimSpace(string(bs))); err != nil {
		t.Fatalf("expect an integer; but got %s", bs)
	} else {
		if num != 6 {
			t.Fatalf("expected to be 6; but got %d", num)
		}
	}
}

func TestRouting(t *testing.T) {
	h := handler()
	srv := httptest.NewServer(h)

	res, err := http.Get(fmt.Sprintf("%s/double?v=4", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request to route double: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expect status OK; but got %d", res.StatusCode)
	}
}
