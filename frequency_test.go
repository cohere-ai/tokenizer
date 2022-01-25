package tokenizer

import (
	"log"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestFrequencySuite(t *testing.T) {
	t.Run("CountString", testCountString)
	t.Run("FrequencyCounter.CountReader", testCountReader)
	t.Run("FrequencyCounter.AddCounts", testAddCounts)
}

func testCountString(t *testing.T) {
	tests := []struct {
		input               string
		expectedFrequencies map[string]int64
	}{
		{
			input: "today my friend and I went running. I liked it.",
			expectedFrequencies: map[string]int64{
				"today":    1,
				"ĠI":       2,
				"Ġmy":      1,
				"Ġfriend":  1,
				"Ġand":     1,
				"Ġwent":    1,
				"Ġrunning": 1,
				".":        2,
				"Ġliked":   1,
				"Ġit":      1,
			},
		},
		{
			input: "起来Qǐlái!！ 不愿Búyuàn做zuò奴隶núlì的de人们rénmen!",
			expectedFrequencies: map[string]int64{
				"!":             1,
				"!ï¼ģ":          1,
				"èµ·æĿ¥QÇĲlÃ¡i": 1,
				"Ġä¸įæĦ¿BÃºyuÃłnåģļzuÃ²å¥´éļ¶nÃºlÃ¬çļĦdeäººä»¬rÃ©nmen": 1,
			},
		},
		{
			input: "a b c d e f g h i j k l 		m n o p q r s t u v w x y z",
			expectedFrequencies: map[string]int64{
				"a":  1,
				"m":  1,
				"ĉ":  1,
				"Ġb": 1,
				"Ġc": 1,
				"Ġd": 1,
				"Ġe": 1,
				"Ġf": 1,
				"Ġg": 1,
				"Ġh": 1,
				"Ġi": 1,
				"Ġj": 1,
				"Ġk": 1,
				"Ġl": 1,
				"Ġn": 1,
				"Ġo": 1,
				"Ġp": 1,
				"Ġq": 1,
				"Ġr": 1,
				"Ġs": 1,
				"Ġt": 1,
				"Ġu": 1,
				"Ġv": 1,
				"Ġw": 1,
				"Ġx": 1,
				"Ġy": 1,
				"Ġz": 1,
				"Ġĉ": 1,
			},
		},
		{
			input: "🐋🐳 🤯",
			expectedFrequencies: map[string]int64{
				"ðŁĲĭðŁĲ³": 1,
				"ĠðŁ¤¯":    1,
			},
		},
	}

	for _, tt := range tests {
		counts := CountString(tt.input)
		if len(counts) != len(tt.expectedFrequencies) {
			t.Fatalf("expected %d words but got %d", len(tt.expectedFrequencies), len(counts))
		}
		for expectedk, expectedv := range tt.expectedFrequencies {
			v, ok := counts[expectedk]
			if !ok {
				t.Fatalf("expected frequencies to contain \"%s\"", expectedk)
			}

			if expectedv != v {
				t.Fatalf("expected %s to have count %d but got %d", expectedk, expectedv, v)
			}
		}
	}
}

func testCountReader(t *testing.T) {
	tests := []struct {
		input               string
		expectedFrequencies map[string]int64
	}{
		{
			input: `today my friend and I went running. I liked it.
			起来Qǐlái!！ 不愿Búyuàn做zuò奴隶núlì的de人们rénmen!
			a b c d e f g h i j k l 		m n o p q r s t u v w x y z
			🐋🐳 🤯
			`,
			expectedFrequencies: map[string]int64{
				"today":         1,
				"ĠI":            2,
				"Ġmy":           1,
				"Ġfriend":       1,
				"Ġand":          1,
				"Ġwent":         1,
				"Ġrunning":      1,
				".":             2,
				"Ġliked":        1,
				"Ġit":           1,
				"!":             1,
				"!ï¼ģ":          1,
				"èµ·æĿ¥QÇĲlÃ¡i": 1,
				"Ġä¸įæĦ¿BÃºyuÃłnåģļzuÃ²å¥´éļ¶nÃºlÃ¬çļĦdeäººä»¬rÃ©nmen": 1,
				"a":        1,
				"m":        1,
				"Ġb":       1,
				"Ġc":       1,
				"Ġd":       1,
				"Ġe":       1,
				"Ġf":       1,
				"Ġg":       1,
				"Ġh":       1,
				"Ġi":       1,
				"Ġj":       1,
				"Ġk":       1,
				"Ġl":       1,
				"Ġn":       1,
				"Ġo":       1,
				"Ġp":       1,
				"Ġq":       1,
				"Ġr":       1,
				"Ġs":       1,
				"Ġt":       1,
				"Ġu":       1,
				"Ġv":       1,
				"Ġw":       1,
				"Ġx":       1,
				"Ġy":       1,
				"Ġz":       1,
				"Ġĉ":       1,
				"ðŁĲĭðŁĲ³": 1,
				"ĠðŁ¤¯":    1,
				"ĉ":        4,
				"ĉĉ":       3,
				"ĉĉĉ":      1,
				"Ċ":        4,
			},
		},
	}

	for _, tt := range tests {
		freq, err := CountReader(strings.NewReader(tt.input))
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to count reader"))
		}

		if len(freq) != len(tt.expectedFrequencies) {
			t.Fatalf("expected %d words but got %d", len(tt.expectedFrequencies), len(freq))
		}

		for expectedk, expectedv := range tt.expectedFrequencies {
			v, ok := freq[expectedk]
			if !ok {
				t.Fatalf("expected frequencies to contain \"%s\"", expectedk)
			}

			if expectedv != v {
				t.Fatalf("expected %s to have count %d but got %d", expectedk, expectedv, v)
			}
		}
	}
}

func testAddCounts(t *testing.T) {
	tests := []struct {
		initial             map[string]int64
		input               map[string]int64
		expectedFrequencies map[string]int64
		expectedNumWords    int64
	}{
		{
			initial: map[string]int64{},
			input: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expectedFrequencies: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
		{
			initial: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			input: nil,
			expectedFrequencies: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
		{
			initial: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			input: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expectedFrequencies: map[string]int64{
				"a": 2,
				"b": 4,
				"c": 6,
			},
		},
		{
			initial: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			input: map[string]int64{
				"d": 1,
				"e": 2,
				"f": 3,
			},
			expectedFrequencies: map[string]int64{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 1,
				"e": 2,
				"f": 3,
			},
		},
	}

	for _, tt := range tests {
		counter := tt.initial
		MergeCounts(counter, tt.input)
		if len(counter) != len(tt.expectedFrequencies) {
			t.Fatalf("expected %d words but got %d", len(tt.expectedFrequencies), len(counter))
		}

		for expectedk, expectedv := range tt.expectedFrequencies {
			v, ok := counter[expectedk]
			if !ok {
				t.Fatalf("expected frequencies to contain \"%s\"", expectedk)
			}

			if expectedv != v {
				t.Fatalf("expected %s to have count %d but got %d", expectedk, expectedv, v)
			}
		}
	}
}
