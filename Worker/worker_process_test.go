package worker

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "test.txt")
	defer tempFile.Close()
	if err != nil {
		t.Fatalf("Failed to Create File")
	}
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	return tempFile.Name()
}

func TestProcessFile(t *testing.T) {
	fileName := createTempFile(t, "hello world\nhello go\n")
	defer os.Remove(fileName)

	lines, wordFreq, err := processFile(fileName)

	expectedMap := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    1,
	}
	expectedLineCount := 2
	//t.Logf("wordFreq: %#v", wordFreq)
	assert.NoError(t, err)
	assert.Equal(t, expectedLineCount, lines)
	assert.Equal(t, expectedMap["hello"], wordFreq["hello"])
	assert.Equal(t, expectedMap["world"], wordFreq["world"])
	assert.Equal(t, expectedMap["go"], wordFreq["go"])

}

func TestProcessFile_Empty(t *testing.T) {
	filename := createTempFile(t, "")
	defer os.Remove(filename)

	lines, wordFreq, err := processFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, 0, lines)
	assert.Equal(t, 0, len(wordFreq))
}
