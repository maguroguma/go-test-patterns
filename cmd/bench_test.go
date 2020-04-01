package main

import "testing"

func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, cat},
		{"Buf", 3, buf},
	}

	for _, c := range benchCases {
		b.Run(
			fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) {
				bench(b, c.n, c.f)
			}
		)
	}
}
