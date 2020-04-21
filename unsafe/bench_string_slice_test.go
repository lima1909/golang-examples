package main

import "testing"

func BenchmarkStingSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := string2bytes("Hello")
		if len(r) <= 0 {
			b.Fatal(len(r))
		}

	}
}

func BenchmarkStingSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := []byte("Hello")
		if len(r) <= 0 {
			b.Fatal(len(r))
		}
	}
}
