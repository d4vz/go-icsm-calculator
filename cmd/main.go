package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/d4vz/go-icms-calculator/internal/icms"
	"github.com/d4vz/go-icms-calculator/internal/random"
	"github.com/d4vz/go-icms-calculator/internal/sell"
	"github.com/d4vz/go-icms-calculator/messaging"
)

var ufs = []string{"SP", "RJ", "MG", "PR", "SC"}

func main() {
	pubSub := messaging.NewPubSub[*sell.Sell](100)
	defer pubSub.Close()
	icmsCalculator := icms.NewICMSCalculator()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for range int(10) {
			pubSub.Publish(&sell.Sell{
				ProductValue:  random.RandomFloat64(100, 1000),
				UfOrigin:      random.RandomInArray(ufs),
				UfDestination: random.RandomInArray(ufs),
			})
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for sell := range pubSub.Subscribe() {
			icmsValue := icmsCalculator.CalculateFor(sell)
			sell.ICMSValue = icmsValue
			fmt.Printf("Sell: %+v\n", sell)
		}
	}()

	<-sigChan
	fmt.Println("Encerrando aplicação...")
}
