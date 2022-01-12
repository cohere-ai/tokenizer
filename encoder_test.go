package tokenizer

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Name string
	Err  error
}

var characterSet = []rune("1234567890[]',.pyfgcrl/=\aoeuidhtns-;qjkxbmwvz!@#$%^&*(){}\"<>PYFGCRL?+|AOEUIDHTNS_:QJKXBMWVZ")

func defaultEncoder(t *testing.T) *Encoder {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	require.NoError(t, err)
	return encoder
}

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

func TestUnicodeEncode(t *testing.T) {
	testCases := []struct {
		testCase   TestCase
		inputWord  string
		outputWord string
	}{
		{
			testCase:   TestCase{Name: "normal word"},
			inputWord:  "testing",
			outputWord: "testing",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testCase.Name, func(tt *testing.T) {
			require.Equal(tt, tc.outputWord, unicodeEncode(tc.inputWord))
		})
	}
}
func TestEncodeDecodeSuccess(t *testing.T) {
	encoder := defaultEncoder(t)

	testCases := []struct {
		testCase TestCase
		tokens   []string
	}{
		{
			testCase: TestCase{Name: "{ }"},
			tokens: []string{
				" ",
			},
		},
		{
			testCase: TestCase{Name: "a"},
			tokens: []string{
				"a",
			},
		},
		{
			testCase: TestCase{Name: "{ }apple"},
			tokens: []string{
				" apple",
			},
		},
		{
			testCase: TestCase{Name: "lorem ipsum"},
			tokens: []string{
				"L", "orem", " ipsum", " dolor", " sit", " amet", ",", " consectetur", " adip", "iscing", " elit", ".", " N", "ulla", " quis", ".",
			},
		},
		{
			testCase: TestCase{Name: "weird character"},
			tokens: []string{
				"È",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testCase.Name, func(tt *testing.T) {
			joinedTokens := strings.Join(tc.tokens, "")
			encoded := encoder.Encode(joinedTokens)

			require.Len(t, encoded, len(tc.tokens))
			for i, token := range tc.tokens {
				require.Equal(t, token, encoder.Decode([]int64{encoded[i]}))
			}

			require.Equal(t, joinedTokens, encoder.Decode(encoded))
		})
	}
}

// benchmarking 1k token speed
func Benchmark1000TokensDecode(b *testing.B) { benchmarkDecode(1000, b) }
func Benchmark1000TokensEncode(b *testing.B) { benchmarkTokenDecode(1000, b) }

func generateTokens(numTokens int) []int64 {
	var tokens []int64
	for n := 0; n < numTokens; n++ {
		tokens = append(tokens, rand.Int63n(50000-1)+1)
	}
	return tokens
}

func benchmarkTokenDecode(numTokens int, b *testing.B) {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	if err != nil {
		b.Error(err)
	}

	tokens := generateTokens(numTokens)
	s := encoder.Decode(tokens)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		encoder.Encode(s)
	}
}

func benchmarkDecode(numTokens int, b *testing.B) {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	if err != nil {
		b.Error(err)
	}
	tokens := generateTokens(numTokens)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encoder.Decode(tokens)
	}
}
