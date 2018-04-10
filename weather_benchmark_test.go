package weather

import "testing"

func Benchmark_toDigital(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toDigital(89)
	}
}
