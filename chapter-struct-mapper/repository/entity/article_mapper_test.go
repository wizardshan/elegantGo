package entity

import (
	"testing"
)

// go test -bench . -benchmem
func BenchmarkMapper(b *testing.B) {
	for n := 0; n < b.N; n++ {
		entArticle := new(Article)
		entArticle.Mapper()
	}
}

func BenchmarkMapperWithCopier(b *testing.B) {
	for n := 0; n < b.N; n++ {
		entArticle := new(Article)
		entArticle.MapperCopier()
	}
}
