package main

import "testing"

func BenchmarkEvil(b *testing.B) {
	f := Filter{
		Name:     "a",
		Lastname: "b",
		Search:   "",
		Limit:    10,
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = QueryBuilderEvil(f)
	}
}

func BenchmarkOK(b *testing.B) {
	f := Filter{
		Name:     "a",
		Lastname: "b",
		Search:   "",
		Limit:    10,
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = QueryBuilderOK(f)
	}
}

func BenchmarkOKAlter(b *testing.B) {
	f := Filter{
		Name:     "a",
		Lastname: "b",
		Search:   "",
		Limit:    10,
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = QueryBuilderOKAlter(f)
	}
}

func BenchmarkOKMask(b *testing.B) {
	f := Filter{
		Name:     "a",
		Lastname: "b",
		Search:   "",
		Limit:    10,
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = QueryBuilderOKMask(f)
	}
}
