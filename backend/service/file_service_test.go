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
	strDir, dirErr := path.StringDir()
	strPath, pathErr := path.StringPath()

	if dirErr != nil || pathErr != nil {
		t.Fail()
	}

	assert.NotEmpty(t, strDir)
	assert.NotEmpty(t, strPath)
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

	strPathAnte, pathErrAnte := path.StringPath()
	if pathErrAnte != nil {
		t.Fail()
	}
	statAnte := fileExistence(strPathAnte)

	// make directories if not exist
	defer os.RemoveAll(path.BaseDir)
	if _, err := os.Stat(path.BaseDir); os.IsNotExist(err) {
		os.Mkdir(path.BaseDir, os.FileMode(0775))
	}
	strDir, dirErr := path.StringDir()
	if dirErr != nil {
		t.Fail()
	}

	if _, err := os.Stat(strDir); os.IsNotExist(err) {
		os.Mkdir(strDir, os.FileMode(0775))
	}

	// WHEN
	err := FileService.CreateFile(path, file)
	fmt.Println(err)

	// THEN
	strPath, pathErr := path.StringPath()
	if pathErr != nil {
		t.Fail()
	}
	statPost := fileExistence(strPath)

	assert.False(t, statAnte)
	assert.True(t, statPost)

	os.Remove(strPath)
}

func fileExistence(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
