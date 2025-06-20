package icms

import "github.com/d4vz/go-icms-calculator/internal/sell"

type ICMSCalculator struct {
}

func NewICMSCalculator() *ICMSCalculator {
	return &ICMSCalculator{}
}

func (c *ICMSCalculator) CalculateFor(sell *sell.Sell) float64 {
	aliquota := getAliquotaICMS(sell.UfOrigin, sell.UfDestination)
	return sell.ProductValue * aliquota / 100
}
