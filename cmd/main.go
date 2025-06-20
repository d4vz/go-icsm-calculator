package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/d4vz/go-icms-calculator/backpressure"
	"github.com/d4vz/go-icms-calculator/internal/config"
	"github.com/d4vz/go-icms-calculator/internal/icms"
	"github.com/d4vz/go-icms-calculator/internal/random"
	"github.com/d4vz/go-icms-calculator/internal/sell"
	"github.com/d4vz/go-icms-calculator/messaging"
)

var ufs = []string{"SP", "RJ", "MG", "PR", "SC"}

func init() {
	config.Load()
}

func main() {
	pubSub := messaging.NewPubSub[*sell.Sell](10)
	defer pubSub.Close()
	icmsCalculator := icms.NewICMSCalculator()
	backpressure := backpressure.NewBackpressure(config.Get().EnableBackpressure, config.Get().BackpressureMaxAttempts)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for range 20 {
			sell := &sell.Sell{
				ProductValue:  random.RandomFloat64(100, 1000),
				UfOrigin:      random.RandomInArray(ufs),
				UfDestination: random.RandomInArray(ufs),
			}

			err := backpressure.WithBackpressure(func() error {
				return pubSub.Publish(sell)
			})

			if err != nil {
				fmt.Printf("Error publishing sell: %v\n", err)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for sell := range pubSub.Subscribe() {
			icmsValue := icmsCalculator.CalculateFor(sell)
			sell.ICMSValue = icmsValue
			fmt.Printf("Sell processed: %+v\n", sell)
			time.Sleep(5 * time.Second)
		}
	}()

	<-sigChan
	fmt.Println("Encerrando aplicação...")
}
