package extract

import (
	"testing"
)

func BenchmarkExtractByMarshalling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtractByMarshalling()
	}
}
func BenchmarkExtractManually(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtractManually()
	}
}
func BenchmarkExtractManuallyPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtractManuallyPointer()
	}
}
