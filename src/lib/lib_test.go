package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBoldFiftyPercentOfWordWithPunctuationInMD(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		TypeBold: "md",
		Percent:  50,
		Jumps:    0,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "**Th**e **su**n **ris**es **i**n **th**e **ea**st.")
}

func TestShouldBoldFiftyPercentOfWordInMD(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		TypeBold: "md",
		Percent:  50,
		Jumps:    0,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "**Th**e **su**n **ris**es **i**n **th**e **ea**st")
}

func TestShouldBoldFiftyPercentOfWordInHTML(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		TypeBold: "html",
		Percent:  50,
		Jumps:    0,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>Th</b>e <b>su</b>n <b>ris</b>es <b>i</b>n <b>th</b>e <b>ea</b>st")
}

func TestShouldBoldFiftyPercentOfWordWithPunctuationInHTML(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		TypeBold: "html",
		Percent:  50,
		Jumps:    0,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>Th</b>e <b>su</b>n <b>ris</b>es <b>i</b>n <b>th</b>e <b>ea</b>st.")
}

func TestShouldBoldFiftyPercentOfMultilineTextInMD(t *testing.T) {
	word := `The sun rises
	early in the morning,
	illuminating the sky
	and announcing
	the start of a new day
	in the east.`
	args := BoldArgs{
		TypeBold: "md",
		Percent:  50,
	}
	boldWord := MakeTextToBold(word, args)
	expected := `**Th**e **su**n **ris**es
	**ear**ly **i**n **th**e **morn**ing,
	**illumi**nating **th**e **sk**y
	**an**d **annou**ncing
	**th**e **sta**rt **o**f **a** **ne**w **da**y
	**i**n **th**e **ea**st.`
	assert.Equal(t, expected, boldWord)
}

func TestShouldBoldVariousPercentagesAboveFiftyInMD(t *testing.T) {
	tests := map[int]struct {
		Input    string
		Expected string
	}{
		55: {
			Input:    "The sun rises in the east.",
			Expected: "**Th**e **su**n **ris**es **i**n **th**e **ea**st.",
		},
		60: {
			Input:    "The sun rises in the east.",
			Expected: "**Th**e **su**n **ris**es **i**n **th**e **ea**st.",
		},
		70: {
			Input:    "The sun rises in the east.",
			Expected: "**Th**e **su**n **rise**s **i**n **th**e **eas**t.",
		},
		80: {
			Input:    "Brightness covers the vast sky.",
			Expected: "**Brightne**ss **cover**s **th**e **vas**t **sk**y.",
		},
		90: {
			Input:    "Shadows stretch across the land.",
			Expected: "**Shadow**s **stretc**h **acros**s **the** **land**.",
		},
		100: {
			Input:    "The sun rises in the east.",
			Expected: "**The** **sun** **rises** **in** **the** **east**.",
		},
	}

	for Percent, tc := range tests {
		args := BoldArgs{
			TypeBold: "md",
			Percent:  Percent,
			Jumps:    0,
		}
		boldWord := MakeTextToBold(tc.Input, args)
		assert.Equal(t, tc.Expected, boldWord, "Failed for percentage: %d", Percent)
	}
}

func TestShouldBoldFiftyPercentWithJumpWords(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		TypeBold: "md",
		Percent:  50,
		Jumps:    1,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "**Th**e sun **ris**es in **th**e east.")
}

func TestShouldBoldFiftyPercentOfWordInANSI(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		TypeBold: "ansi",
		Percent:  50,
		Jumps:    0,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, "\x1b[1mTh\x1b[0me \x1b[1msu\x1b[0mn \x1b[1mris\x1b[0mes \x1b[1mi\x1b[0mn \x1b[1mth\x1b[0me \x1b[1mea\x1b[0mst", boldWord)
}

func TestBionicCurveByLength(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"I", "I"},                               // len 1 → no bold
		{"to", "**t**o"},                         // len 2 → 1
		{"sun", "**s**un"},                       // len 3 → 1
		{"east", "**ea**st"},                     // len 4 → 2
		{"rises", "**ris**es"},                   // len 5 → 3
		{"bright", "**bri**ght"},                 // len 6 → 3
		{"morning", "**morn**ing"},               // len 7 → 4
		{"shadowed", "**shad**owed"},             // len 8 → 4
		{"announces", "**anno**unces"},           // len 9 → 4
		{"illuminate", "**illu**minate"},         // len 10 → int(4.5) = 4
		{"transformation", "**transf**ormation"}, // len 14 → round(5.6) = 6
	}
	args := BoldArgs{TypeBold: "md", Percent: 0, Jumps: 0}
	for _, tc := range cases {
		out := MakeTextToBold(tc.input, args)
		assert.Equal(t, tc.expected, out, "input: %q", tc.input)
	}
}

func TestBionicCurveOnSentence(t *testing.T) {
	input := "The sun rises in the east."
	args := BoldArgs{TypeBold: "md", Percent: 0, Jumps: 0}
	out := MakeTextToBold(input, args)
	// The(3)→1, sun(3)→1, rises(5)→3, in(2)→1, the(3)→1, east(4)→2
	assert.Equal(t, "**T**he **s**un **ris**es **i**n **t**he **ea**st.", out)
}

func TestBionicCurveSkipsSingleLetterWord(t *testing.T) {
	input := "I am here"
	args := BoldArgs{TypeBold: "md", Percent: 0, Jumps: 0}
	out := MakeTextToBold(input, args)
	// I(1)→0, am(2)→1, here(4)→2
	assert.Equal(t, "I **a**m **he**re", out)
}

func TestUnicodeWords(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 0, Jumps: 0}
	out := MakeTextToBold("café não está aqui", args)
	// café(4)→2, não(3)→1, está(4)→2, aqui(4)→2
	assert.Equal(t, "**ca**fé **n**ão **es**tá **aq**ui", out)
}

func TestEdgeCaseEmpty(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 50, Jumps: 0}
	assert.Equal(t, "", MakeTextToBold("", args))
}

func TestEdgeCaseWhitespaceOnly(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 50, Jumps: 0}
	assert.Equal(t, "   \t  ", MakeTextToBold("   \t  ", args))
}

func TestEdgeCasePunctuationOnly(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 50, Jumps: 0}
	assert.Equal(t, "...!?,;", MakeTextToBold("...!?,;", args))
}

func TestEdgeCasePreservesEmbeddedNewlines(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 50, Jumps: 0}
	out := MakeTextToBold("hi\nthere", args)
	// "hi" len 2 Percent 50 → int(1+0.5) = 1 → "**h**i"
	// "there" len 5 Percent 50 → int(2.5+0.5) = 3 → "**the**re"
	assert.Equal(t, "**h**i\n**the**re", out)
}

func TestEdgeCaseDigitsAsWordBoundary(t *testing.T) {
	args := BoldArgs{TypeBold: "md", Percent: 0, Jumps: 0}
	// "covid19" with curve: "covid"(5)→3 bold, "19" digits act as boundary (no markers)
	assert.Equal(t, "**cov**id19", MakeTextToBold("covid19", args))
	// "abc1" used to lose the "abc" entirely; should now emit "abc" bolded then "1"
	assert.Equal(t, "**a**bc1", MakeTextToBold("abc1", args))
	// Standalone digits: "5 apples" → "5"(1)→0 bold, then "apples"(6)→3 bold
	assert.Equal(t, "5 **app**les", MakeTextToBold("5 apples", args))
}
