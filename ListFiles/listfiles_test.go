package listfiles

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFile(t *testing.T) {
	dir := t.TempDir()

	txt1 := filepath.Join(dir, "file11.txt")
	txt2 := filepath.Join(dir, "file12.txt")
	nonTxt := filepath.Join(dir, "file13.log")
	os.WriteFile(txt1, []byte("hello"), 0644)
	os.WriteFile(txt2, []byte("world"), 0644)
	os.WriteFile(nonTxt, []byte("nope"), 0644)

	files, err := ListFile(dir)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{txt1, txt2}, files)
}

func TestListFile_DirNotExists(t *testing.T) {
	nonExistentDir := filepath.Join(os.TempDir(), "Dir1")
	files, err := ListFile(nonExistentDir)
	assert.Error(t, err)
	assert.Nil(t, files)
}
