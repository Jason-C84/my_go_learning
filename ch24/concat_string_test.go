package ch24

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByBuf(b *testing.B) {
	strs := []string{"1", "2", "3", "4", "5"}
	var ret bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range strs {
			ret.WriteString(v)
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	strs := []string{"1", "2", "3", "4", "5"}
	var ret string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range strs {
			ret += v
		}

	}
	b.StopTimer()
}
