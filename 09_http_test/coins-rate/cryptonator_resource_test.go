package coins_rate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCryptonatorResource_BitCoinToUSDRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ticker":{"base":"BTC","target":"USD","price":"4007.93579679","volume":"32581.54482720","change":"2.52534325"},"timestamp":1552993382,"success":true,"error":""}`))
		return
	}))

	resource := cryptonatorResource{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err != nil {
		t.Error(err)
	}

	if result == 0 {
		t.Fail()
	}
}

//TODO other tests
