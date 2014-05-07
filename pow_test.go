package pow

import "testing"

func BenchmarkDistanceWithPow(b *testing.B) {
	v1 := NewVector(1000, 0, 5000)
	v2 := NewVector(1000, 0, 5000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DistanceWithPow(v1, v2)
	}
}

func BenchmarkDistanceWithoutPow(b *testing.B) {
	v1 := NewVector(1000, 0, 5000)
	v2 := NewVector(1000, 0, 5000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DistanceWithoutPow(v1, v2)
	}
}
