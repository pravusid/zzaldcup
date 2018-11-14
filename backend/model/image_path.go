package model

import (
	"errors"
	"os"
	"strings"
)

type ImagePath struct {
	BaseDir   string `json:"baseDir"`
	Shard     string `json:"shard"`
	Checksum  string `json:"checksum"`
	Extension string `json:"extension"`
}

// StringDir returns path without filename and extension
func (img *ImagePath) StringDir() (string, error) {
	if len(img.Shard) == 0 {
		return "", errors.New("ImagePath: shard is not exist")
	}

	dir := strings.Join([]string{img.BaseDir, img.Shard}, string(os.PathSeparator))
	if img.BaseDir == "" {
		dir = strings.TrimPrefix(dir, string(os.PathSeparator))
	}
	return dir + string(os.PathSeparator), nil
}

// StringPath returns entire path of the file
func (img *ImagePath) StringPath() (string, error) {
	if len(img.Checksum) == 0 {
		return "", errors.New("ImagePath: checksum is not exist")
	}

	dir, err := img.StringDir()
	if err != nil {
		return "", err
	}

	return dir + img.Checksum + img.Extension, nil
}
