package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouting_RateBTC(t *testing.T) {
	srv := httptest.NewServer(handlers())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/rate/btc", srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "BitCoin to USD rate: 0.000000 $\n" {
		t.Fail()
	}

}
