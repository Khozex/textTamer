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
