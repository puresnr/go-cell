package cast

import (
	"fmt"
	"github.com/puresnr/go-cell/generic"
	"strconv"
	"testing"
)

var num = 50000

func toString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func toStringFg[T1 generic.Signed](i T1) string {
	return fmt.Sprint(i)
}

func toStringF(i interface{}) string {
	return fmt.Sprint(i)
}

func toStringFdg[T1 generic.Signed](i T1) string {
	return fmt.Sprintf("%d", i)
}

func toStringFd(i interface{}) string {
	return fmt.Sprintf("%d", i)
}

func toStringFv(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func toStringFvg[T1 generic.Signed](i T1) string {
	return fmt.Sprintf("%v", i)
}

func BenchmarkToStringI(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toString(num)
	}
}

func BenchmarkToString(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToString(num)
	}
}

func BenchmarkToStringFg(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringFg(num)
	}
}

func BenchmarkToStringF(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringF(num)
	}
}

func BenchmarkToStringFdg(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringFdg(num)
	}
}

func BenchmarkToStringFd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringFd(num)
	}
}

func BenchmarkToStringFv(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringFv(num)
	}
}

func BenchmarkToStringFvg(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toStringFvg(num)
	}
}
