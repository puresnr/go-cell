package datetime

import "testing"

func BenchmarkLawAgeNow(b *testing.B) {
	for i := 0; i != b.N; i++ {
		LawAgeNow("1988-12-03")
	}
}
