package pow

import (
	"math"
	"math/rand"
)

type Vector []float64

func NewVector(size, minValue, maxValue int) Vector {
	v := make(Vector, size)
	width := float64(maxValue - minValue)
	for i := range v {
		v[i] = rand.Float64()*width + float64(minValue)
	}
	return v
}

func DistanceWithPow(v1, v2 Vector) float64 {
	// assumes that len(v1) == len(v2)
	d := 0.0
	for i := range v1 {
		d += math.Pow(v1[i]-v2[i], 2)
	}
	return d
}

func DistanceWithoutPow(v1, v2 Vector) float64 {
	// also assumes that len(v1) == len(v2)
	d := 0.0
	for i := range v1 {
		temp := v1[i] - v2[i]
		d += temp * temp
	}
	return d
}
