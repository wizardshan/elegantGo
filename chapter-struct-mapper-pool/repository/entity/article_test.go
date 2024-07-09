package entity

import (
	"testing"
)

// go test -bench . -benchmem
func BenchmarkMapper(b *testing.B) {
	var entArticle Article
	for n := 0; n < b.N; n++ {
		entArticle.Mapper()
	}
}

func BenchmarkMapperWithCopier(b *testing.B) {
	var entArticle Article
	for n := 0; n < b.N; n++ {
		entArticle.MapperCopier()
	}
}

func BenchmarkMapperWithPool(b *testing.B) {
	var entArticle Article
	for n := 0; n < b.N; n++ {
		entArticle.MapperPool()
	}
}
