package main

import "testing"

// go test -bench . -benchmem
func BenchmarkArticle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Set()
	}
}
