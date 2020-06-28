package main

import (
	"math"
)

// Kubus ...
type Kubus struct {
	Sisi float64
}

// Volume  ...
func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

//Luas ...Luas
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

//Keliling ...Keliling
func (k Kubus) Keliling() float64 {
	return k.Sisi * 4
}
