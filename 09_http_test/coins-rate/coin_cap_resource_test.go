package coins_rate

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCoinCapResource_BitCoinToUSDRate(t *testing.T) {
	// create a new server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// returns JSON regarding to the docs
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"data":{"id":"bitcoin","symbol":"BTC",
		"currencySymbol":"₿","type":"crypto","rateUsd":"4010.8714336221081818"},"timestamp":1552990697033}
		`))
		return
	}))

	resource := coinCapResource{
		httpClient: srv.Client(), // test http Client
		baseURL:    srv.URL,      // test server URL
	}

	// make the request
	res, err := resource.BitCoinToUSDRate(nil)
	if err != nil {
		t.Error(err)
	}
	// if the res is equal to 0
	if res == 0 {
		t.Fail()
	}
}

func TestCoinCapResource_BitCoinToUSDRateNotOK(t *testing.T) {
	// create a new server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// returns JSON regarding to the docs
		w.WriteHeader(http.StatusBadRequest)

		return
	}))

	resource := coinCapResource{
		httpClient: srv.Client(), // test http Client
		baseURL:    srv.URL,      // test server URL
	}

	// make the request
	res, err := resource.BitCoinToUSDRate(nil)
	if err == nil {
		t.Fail()
	}
	// if the error text is wrong
	if err.Error() != ErrTextNotOK {
		t.Fail()
	}

	// We wait the err, so if the res is not 0 -> Fail
	if res != 0 {
		t.Fail()
	}
}

func TestCoinCapResource_BitCoinToUSDRateTimout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// hold the responce to make simulate timeout
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusNoContent)
		return
	}))

	resource := coinCapResource{
		httpClient: srv.Client(), // test http Client
		baseURL:    srv.URL,      // test server URL
	}

	// make the request
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Millisecond)
	res, err := resource.BitCoinToUSDRate(ctx)
	// we need the errors, otherwise -> fail the test
	if err == nil {
		t.Fail()
	}

	// We wait the err, so if the res is not 0 -> Fail
	if res != 0 {
		t.Fail()
	}

}
