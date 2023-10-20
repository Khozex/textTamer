package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAFivePercentWordBoldWithPunctuation(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		typeBold: "md",
		porcent:  50,
	}
	boldWord := makeTextToBold(word, args)
	assert.Equal(t, boldWord, "**T**he **s**un **ri**ses **i**n **t**he **ea**st.")
}

func TestMakeAFivePercentWordBold(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		typeBold: "md",
		porcent:  50,
	}
	boldWord := makeTextToBold(word, args)
	assert.Equal(t, boldWord, "**T**he **s**un **ri**ses **i**n **t**he **ea**st")
}

func TestMakeAFivePercentWordBoldWithHtml(t *testing.T) {
	word := "The sun rises in the east"
	args := BoldArgs{
		typeBold: "html",
		porcent:  50,
	}
	boldWord := makeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>T</b>he <b>s</b>un <b>ri</b>ses <b>i</b>n <b>t</b>he <b>ea</b>st")
}

func TestMakeAFivePercentWordBoldWithHtmlAndPunctuation(t *testing.T) {
	word := "The sun rises in the east."
	args := BoldArgs{
		typeBold: "html",
		porcent:  50,
	}
	boldWord := makeTextToBold(word, args)
	assert.Equal(t, boldWord, "<b>T</b>he <b>s</b>un <b>ri</b>ses <b>i</b>n <b>t</b>he <b>ea</b>st.")
}

func TestMakeWithBreakTexts(t *testing.T) {
	word := `The sun rises
	early in the morning,
	illuminating the sky
	and announcing
	the start of a new day
	in the east.`
	args := BoldArgs{
		typeBold: "md",
		porcent:  50,
	}
	boldWord := makeTextToBold(word, args)
	assert.Equal(t, boldWord, `**T**he **s**un **ri**ses
	**ea**rly **i**n **t**he **mor**ning,
	**illumi**nating **t**he **s**ky
	**a**nd **annou**ncing
	**t**he **st**art **o**f a **n**ew **d**ay
	**i**n **t**he **ea**st.`)
}
