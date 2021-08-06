package demo_data_gen

import (
	"math/rand"
	"time"
)

func RandomTimeToday() time.Time {
	now := time.Now().Unix()
	yesterday := time.Now().AddDate(0, 0, -1).Unix()
	delta := now - yesterday

	sec := rand.Int63n(delta) + yesterday
	return time.Unix(sec, 0)
}

func RandomFloat64FromRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomString(interfaces []string) string {
	randIndex := rand.Intn(len(interfaces))
	return interfaces[randIndex]
}

func RandomStringIndex(startIndex int, endIndex int, interfaces []string) string {
	randIndex := rand.Intn(endIndex - startIndex)
	return interfaces[startIndex+randIndex]
}
