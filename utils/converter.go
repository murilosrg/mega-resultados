package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func ConvertToInt(text string) int {
	v, _ := strconv.Atoi(text)
	return v
}

func ConvertToDecimal(text string) float64 {
	s := normalize(text)
	d, _ := decimal.NewFromString(s)
	v, _ := d.Float64()
	return v
}

func ConvertToDate(text string) time.Time {
	v, _ := time.Parse("02/01/2006", text)
	return v
}

func ConvertToIntSlice(text ...string) []int {
	var v []int

	for _, t := range text {
		number := ConvertToInt(t)
		v = append(v, number)
	}

	return v
}

func ConvertToBool(text string) bool {
	if text == "SIM" {
		return true
	}
	return false
}

func normalize(old string) string {
	s := strings.Replace(old, ".", "", -1)
	return strings.Replace(s, ",", ".", -1)
}
