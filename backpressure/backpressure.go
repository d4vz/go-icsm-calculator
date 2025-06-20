package backpressure

import (
	"fmt"
	"math"
	"time"
)

type Backpressure struct {
	Enable      bool
	MaxAttempts int
}

func NewBackpressure(enable bool, maxAttempts int) *Backpressure {
	return &Backpressure{
		Enable:      enable,
		MaxAttempts: maxAttempts,
	}
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * 1 * time.Second
}

func (b *Backpressure) WithBackpressure(fn func() error) error {
	if !b.Enable {
		return fn()
	}

	for i := 0; i < b.MaxAttempts; i++ {
		err := fn()

		if err == nil {
			return nil
		}

		fmt.Printf("Error: %v\n", err)
		time.Sleep(exponentialBackoff(i))
	}

	return fmt.Errorf("max attempts reached")
}
