package coins_rate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	ErrTextCommon = "[CoinCap]"
	ErrTextNotOK  = "[CoinCap] not OK status code"
)

type coinCapResource struct {
	httpClient *http.Client
	baseURL    string
}

// return current BitCoin to USD rate on https://coincap.io/
// see https://docs.coincap.io/#c7925509-73f4-4b11-a602-74d1fee44dba
func (rcv *coinCapResource) BitCoinToUSDRate(ctx context.Context) (float64, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/rates/bitcoin", rcv.baseURL), nil)

	if err != nil {
		return 0, errors.Wrap(err, ErrTextCommon)
	}

	// Добавляем заголовок, указывающий какой формат данных мы ожидаем
	r.Header.Set("Accept", "application/json")

	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// Выполняем http запрос
	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return 0, errors.Wrap(err, ErrTextCommon)
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New(ErrTextNotOK)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	var data coinCapResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return 0, errors.Wrap(err, ErrTextCommon)
	}
	return data.Data.RateUsd, err

}

type coinCapResponse struct {
	Data struct {
		// опция string говорит о том, что float64 нужно парсить из строки
		RateUsd float64 `json:"rateUsd,string"`
	} `json:"data"`
}

func NewCoinCapResource(httpClient *http.Client) Resource {
	return &coinCapResource{httpClient: httpClient, baseURL: "https://api.coincap.io"}
}
