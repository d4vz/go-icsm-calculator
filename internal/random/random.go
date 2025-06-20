package random

import "math/rand"

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInArray[T any](array []T) T {
	return array[rand.Intn(len(array))]
}
