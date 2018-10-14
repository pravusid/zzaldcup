package service

import (
	"bytes"
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestFileService_HashingAndBuffering(t *testing.T) {
	// GIVEN
	var file io.Reader = bytes.NewBufferString("something")
	var buffer bytes.Buffer

	// WHEN
	hash, _ := FileService.HashingAndBuffering(&file, &buffer)

	// THEN
	hasher := sha256.New()
	io.Copy(hasher, &buffer)
	assert.Equal(t, hash, hasher.Sum(nil))
}

func TestCompetitorService_GenerateFilePath(t *testing.T) {
	// GIVEN
	hasher := sha256.New()
	var file io.Reader = bytes.NewBufferString("something")
	io.Copy(hasher, file)
	hash := hasher.Sum(nil)

	// WHEN
	path, presence := FileService.GenerateFilePath(hash, "image", ".jpg")

	// THEN
	assert.NotEmpty(t, path)
	assert.False(t, presence)
}

func TestCompetitorService_CreateFile(t *testing.T) {
	// GIVEN
	filename := "test_file"
	var file io.Reader = bytes.NewBufferString("something")
	statAnte := fileExistence(filename)

	// WHEN
	FileService.CreateFile(filename, file)

	// THEN
	statPost := fileExistence(filename)

	assert.False(t, statAnte)
	assert.True(t, statPost)

	os.Remove(filename)
}

func fileExistence(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
