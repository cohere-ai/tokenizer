package tokenizer

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/pkg/errors"
)

var characterSet = []rune("1234567890[]',.pyfgcrl/=\aoeuidhtns-;qjkxbmwvz!@#$%^&*(){}\"<>PYFGCRL?+|AOEUIDHTNS_:QJKXBMWVZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterSet[rand.Intn(len(characterSet))]
	}
	return string(b)
}

func encodeAndDecode(text string, t *testing.T) { // maybe not the best idea, can remove, -- randomized test :/
	encoder, err := NewFromPrebuilt("coheretext-50k")
	if err != nil {
		t.Error(err)
	}

	encoded := encoder.Encode(text)
	decoded := encoder.Decode(encoded)
	if decoded != text {
		err := errors.New("decoded text does not equal input. String that caused this error: " + text)
		t.Error(err)
	}
}

func testSingleTokenEncode(t *testing.T) {
	encoder, err := NewFromPrebuilt("coheretext-50k")
	if err != nil {
		t.Error(err)
	}

	for i := 1; i <= 256; i++ { // back to 50k eventually? to test all values>
		s := []int64{int64(i)}
		if s[0] != encoder.Encode(encoder.Decode(s))[0] {

			err_name := fmt.Sprintf("encoded number is not correct, Output: %d. Expected: %d.", encoder.Encode(encoder.Decode(s))[0], s[0])
			t.Error(errors.New(err_name))
		}
	}
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

/* // in progress function to test both constructors -- is this neccessary?
func testFromPrebuiltAndFromReader(text string, b * testing.B) {
	encoderFromPrebuilt, err1 := NewFromPrebuilt("coheretext-50k")
	if err1 != nil {
		b.Error(err1)
	}

	encoderFromReader, err1 := NewFromReader
}*/

func encodeAndDecodeHelper(num_tests int, t *testing.T) {
	for i := 0; i < num_tests; i++ {
		encodeAndDecode(randomString(100), t)
	}
}

func TestEncodeSingleTokens(t *testing.T) { // does not pass currently -- possible issue w/ special characters?
	testSingleTokenEncode(t)
}

func TestEncodeAndDecode100(t *testing.T)    { encodeAndDecodeHelper(100, t) }
func Benchmark1000TokensDecode(b *testing.B) { benchmarkDecode(1000, b) }
func Benchmark1000TokensEncode(b *testing.B) { benchmarkTokenDecode(1000, b) }
func BenchmarkEncode1Sentence(b *testing.B)  { benchmarkEncode(randomString(100), b) }
func BenchmarkEncode1Paragraph(b *testing.B) { benchmarkEncode(randomString(600), b) }
func BenchmarkEncode1KB(b *testing.B)        { benchmarkEncode(randomString(1000), b) }
func BenchmarkEncode1MB(b *testing.B)        { benchmarkEncode(randomString(1000000), b) }
func BenchmarkEncode500MB(b *testing.B)      { benchmarkEncode(randomString(500000000), b) }
func BenchmarkEncode1GB(b *testing.B)        { benchmarkEncode(randomString(1000000000), b) }
