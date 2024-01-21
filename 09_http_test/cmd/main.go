package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	coins_rate "github.com/alexapps/examples/09_http_test/coins-rate"
)

const (
	defaultPortAddress = ":8080"
	apiAddressKey      = "API_ADDRESS"
)

var (
	apiAddress    string
	httpClient    *http.Client
	rateResources []coins_rate.Resource
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	apiAddress = getVar(apiAddressKey, defaultPortAddress)

	httpClient = &http.Client{}

	// do not make any dependency injection, just initiate the resources
	rateResources := make([]coins_rate.Resource, 3)

	rateResources[0] = coins_rate.NewCoinCapResource(httpClient)
	rateResources[1] = coins_rate.NewCryptonatorResource(httpClient)
	rateResources[2] = coins_rate.NewCryptoCompareResource(httpClient)

}

func main() {
	log.Println(fmt.Sprintf("Server address: %s", apiAddress))
	log.Fatal(http.ListenAndServe(apiAddress, handlers()))

}

func handlers() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("rate/btc", getBitcoinRateHandler)

	return r
}

func getBitcoinRateHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	// if _, err := fmt.Fprintf(w, "BitCoin to USD rate: %f $\n", 0.0); err != nil {
	// 	log.Println(err)
	// }

	var wg sync.WaitGroup
	// to prevent data race with the result variable
	var mux sync.RWMutex
	// will hold te result from go
	var result []float64

	wg.Add(len(rateResources))

	for _, res := range rateResources {
		go func(res coins_rate.Resource) {
			defer wg.Done()

			rate, err := res.BitCoinToUSDRate(nil)

			if err != nil {
				log.Println(err)
				return
			}
			mux.Lock()
			result = append(result, rate)
			mux.Unlock()
		}(res)
	}

	wg.Wait()

	if len(result) != 0 {
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, "There is not result\n"); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprintf(w, "BitCoin to USD rate: %f $\n", avg(result)); err != nil {
		log.Println(err)
	}
}

// some util functions
// getVar returns ENV value by key or default value
func getVar(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// average array numbers
func avg(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var total float64 = 0
	for _, v := range data {
		total += v
	}
	return total / float64(len(data))
}
