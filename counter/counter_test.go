package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergerWordCount(t *testing.T) {
	all := map[string]int{"hello": 2, "world": 1}
	file := map[string]int{"hello": 1, "go": 3}
	MergerWordCount(all, file)
	expected := map[string]int{"hello": 3, "world": 1, "go": 3}
	assert.Equal(t, expected, all)
}

func TestPrintTopWords(t *testing.T) {
	wordFreq := map[string]int{
		"hello": 5,
		"go":    10,
		"world": 7,
		"foo":   2,
	}
	top := PrintTopWords(wordFreq, 2)
	assert.Len(t, top, 2)
	assert.Equal(t, "go", top[0].Word)
	assert.Equal(t, 10, top[0].WordFreq)
	assert.Equal(t, "world", top[1].Word)
	assert.Equal(t, 7, top[1].WordFreq)
}

func TestPrintTopWords_LessThanN(t *testing.T) {
	wordFreq := map[string]int{
		"a": 1,
		"b": 2,
	}
	top := PrintTopWords(wordFreq, 5)
	assert.Len(t, top, 2)
}
