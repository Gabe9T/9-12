package main

import (
	"fmt"
	"math"
)

// Histo struct to hold the histogram buckets
type Histo struct {
	buckets [10]int
}

// Constructor for the Histo struct
func NewHisto(ages []int) *Histo {
	h := &Histo{}
	
	// Fill the buckets based on the ageS
	for _, age := range ages {
		bucketIndex := age / 10
		h.buckets[bucketIndex]++
	}
	
	return h
}

// Method to calculate and return the normalize histogram
func (h *Histo) GetHisto() []float64 {
	total := 0

	// Calculate the total number of ages
	for _, count := range h.buckets {
		total += count
	}

	// Normalize the buckets to percentage
	normalized := make([]float64, 10)
	for i, count := range h.buckets {
		percentage := (float64(count) / float64(total)) * 100
		normalized[i] = math.Round(percentage*100) / 100  // Truncate to 2 decimal places
	}

	return normalized
}

func main() {
	// test program
	ages := []int{1, 67, 99, 21, 55, 87, 23, 33, 11}
	histo := NewHisto(ages)
	result := histo.GetHisto()

	fmt.Println(result)
	
}
