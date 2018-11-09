package service

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-server/model"
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
	assert.NotEmpty(t, path.StringDir())
	assert.NotEmpty(t, path.StringPath())
	assert.False(t, presence)
}

func TestCompetitorService_CreateFile(t *testing.T) {
	// GIVEN
	path := &model.ImagePath{
		BaseDir:   "image",
		Shard:     "0",
		Checksum:  "1234",
		Extension: ".jpg",
	}
	var file io.Reader = bytes.NewBufferString("something")
	statAnte := fileExistence(path.StringPath())

	// make directories if not exist
	if _, err := os.Stat(path.BaseDir); os.IsNotExist(err) {
		os.Mkdir(path.BaseDir, os.FileMode(0775))
	}
	dir := path.StringDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.FileMode(0775))
	}

	// WHEN
	err := FileService.CreateFile(path, file)
	fmt.Println(err)

	// THEN
	statPost := fileExistence(path.StringPath())

	assert.False(t, statAnte)
	assert.True(t, statPost)

	os.Remove(path.StringPath())
}

func fileExistence(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
