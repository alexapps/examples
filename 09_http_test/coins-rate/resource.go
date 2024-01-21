package coins_rate

import "context"

type Resource interface {
	BitCoinToUSDRate(ctx context.Context) (float64, error)
}
