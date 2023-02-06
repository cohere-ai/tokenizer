package tokenizer

import (
	"math/rand"
	"os"
	"reflect"
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

func defaultBenchmarkEncoder(b *testing.B) *Encoder {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	require.NoError(b, err)
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
	b.ReportAllocs()
	encoder := defaultBenchmarkEncoder(b)
	for n := 0; n < b.N; n++ {
		encoder.Encode(text)
	}
}
func BenchmarkEncode1Sentence(b *testing.B)  { benchmarkEncode(randomString(100), b) }
func BenchmarkEncode1Paragraph(b *testing.B) { benchmarkEncode(randomString(600), b) }
func BenchmarkEncode1KB(b *testing.B)        { benchmarkEncode(randomString(1000), b) }
func BenchmarkEncode1MB(b *testing.B)        { benchmarkEncode(randomString(1000000), b) }

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
		{
			testCase: TestCase{Name: "upper-case quotes"},
			tokens: []string{
				"O", "'", "SH", "E", "AN", " DON", "'", "T", " BE", " BA", "'", "D", " '", "MAN", " YOU", "'", "RE", " CO", "ULD", "'", "VE", " HE", "'", "L", "LP", "ED",
			},
		},
		{
			testCase: TestCase{Name: "lower-case quotes"},
			tokens: []string{
				"o", "'s", "he", "an", " don", "'t", " be", " ba", "'d", " '", "man", " you", "'re", " could", "'ve", " he", "'ll", "ped",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testCase.Name, func(tt *testing.T) {
			joinedTokens := strings.Join(tc.tokens, "")
			tokenIDs, tokenStrings := encoder.Encode(joinedTokens)

			for i, token := range tc.tokens {
				require.Equal(t, token, encoder.Decode([]int64{tokenIDs[i]}))
				require.Equal(t, token, tokenStrings[i])
			}

			require.Equal(t, joinedTokens, encoder.Decode(tokenIDs))
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
	b.ReportAllocs()
	encoder := defaultBenchmarkEncoder(b)
	tokens := generateTokens(numTokens)
	s := encoder.Decode(tokens)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		encoder.Encode(s)
	}
}

func benchmarkDecode(numTokens int, b *testing.B) {
	b.ReportAllocs()
	encoder := defaultBenchmarkEncoder(b)
	tokens := generateTokens(numTokens)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encoder.Decode(tokens)
	}
}

func TestFromPrebuiltAndFromReader(t *testing.T) {
	ePrebuilt := defaultEncoder(t)

	encoderReader, err := os.Open("vocab/coheretext-50k/encoder.json")
	require.NoError(t, err)
	vocabReader, err := os.Open("vocab/coheretext-50k/vocab.bpe")
	require.NoError(t, err)

	eReader, err := NewFromReaders(encoderReader, vocabReader)
	require.NoError(t, err)

	if !(reflect.DeepEqual(ePrebuilt.Encoder, eReader.Encoder) &&
		reflect.DeepEqual(ePrebuilt.Decoder, eReader.Decoder) &&
		reflect.DeepEqual(ePrebuilt.BPERanks, eReader.BPERanks) &&
		reflect.DeepEqual(ePrebuilt.Cache, eReader.Cache) &&
		ePrebuilt.VocabSize == eReader.VocabSize) {

		t.Logf("The encoders are not the same.")
		t.Fail()
	}
}
