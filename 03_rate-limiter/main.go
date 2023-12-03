package main

import (
	"sync"
	"time"
)

type Limiter struct {
	mu sync.Mutex
	// Bucket is filled with rate tokens per second
	rate int
	// Bucket size
	bucketSize int
	// Number of tokens in bucket
	nTokens int
	// Time last token was generated
	lastToken time.Time
}

func NewLimiter(rate, limit int) *Limiter {
	return &Limiter{
		rate:       rate,
		bucketSize: limit,
		nTokens:    limit,
		lastToken:  time.Now(),
	}
}

func (s *Limiter) Wait() {
	s.mu.Lock()
	defer s.mu.Unlock()
	// when there are tokens in the bucket, so we simply grab one and return
	if s.nTokens > 0 {
		s.nTokens--
		return
	}

	// Here, there is not enough tokens in the bucket
	tElapsed := time.Since(s.lastToken)
	period := time.Second / time.Duration(s.rate)
	nTokens := tElapsed.Nanoseconds() / period.Nanoseconds()
	s.nTokens = int(nTokens)

	if s.nTokens > s.bucketSize {
		s.nTokens = s.bucketSize
	}
	s.lastToken = s.lastToken.Add(time.Duration(nTokens) * period)
	// We filled the bucket. There may not be enough
	if s.nTokens > 0 {
		s.nTokens--
		return
	}

	// We have to wait until more tokens are available
	// A token should be available at:
	next := s.lastToken.Add(period)
	wait := next.Sub(time.Now())
	if wait >= 0 {
		time.Sleep(wait)
	}
	s.lastToken = next

}

func main() {

}
