package main

import "testing"

func BenchmarkFindPrimes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for prime := range findPrimes(10000) {
			_ = prime
		}
	}
}
