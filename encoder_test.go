package tokenizer

import (
	"math/rand"
	"testing"
)

var characterSet = []rune("1234567890[]',.pyfgcrl/=\aoeuidhtns-;qjkxbmwvz!@#$%^&*(){}\"<>PYFGCRL?+|AOEUIDHTNS_:QJKXBMWVZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterSet[rand.Intn(len(characterSet))]
	}
	return string(b)
}
func benchmarkEncode(text string, b *testing.B) {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		encoder.Encode(text)
	}
}
func BenchmarkEncode1Sentence(b *testing.B)  { benchmarkEncode(randomString(100), b) }
func BenchmarkEncode1Paragraph(b *testing.B) { benchmarkEncode(randomString(600), b) }
func BenchmarkEncode1KB(b *testing.B)        { benchmarkEncode(randomString(1000), b) }
func BenchmarkEncode1MB(b *testing.B)        { benchmarkEncode(randomString(1000000), b) }
func BenchmarkEncode500MB(b *testing.B)      { benchmarkEncode(randomString(500000000), b) }
func BenchmarkEncode1GB(b *testing.B)        { benchmarkEncode(randomString(1000000000), b) }
