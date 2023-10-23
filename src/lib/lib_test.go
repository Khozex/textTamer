package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAFiftyPercentWordBoldWithPunctuation(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		TypeBold: "md",
		Porcent:  50,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "**Th**e **su**n **ris**es **i**n **th**e **ea**st.")
}

func TestMakeAFivePercentWordBold(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		TypeBold: "md",
		Porcent:  50,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "**Th**e **su**n **ris**es **i**n **th**e **ea**st")
}

func TestMakeAFivePercentWordBoldWithHtml(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		TypeBold: "html",
		Porcent:  50,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>Th</b>e <b>su</b>n <b>ris</b>es <b>i</b>n <b>th</b>e <b>ea</b>st")
}

func TestMakeAFivePercentWordBoldWithHtmlAndPunctuation(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		TypeBold: "html",
		Porcent:  50,
	}
	boldWord := MakeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>Th</b>e <b>su</b>n <b>ris</b>es <b>i</b>n <b>th</b>e <b>ea</b>st.")
}

func TestMakeWithBreakTexts(t *testing.T) {
	word := `The sun rises
	early in the morning,
	illuminating the sky
	and announcing
	the start of a new day
	in the east.`
	args := BoldArgs{
		TypeBold: "md",
		Porcent:  50,
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
