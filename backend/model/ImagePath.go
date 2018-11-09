package model

import (
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
func (img *ImagePath) StringDir() string {
	if len(img.Shard) == 0 {
		return ""
	}

	dir := strings.Join([]string{img.BaseDir, img.Shard}, string(os.PathSeparator))
	if img.BaseDir == "" {
		dir = strings.TrimPrefix(dir, string(os.PathSeparator))
	}
	return dir + string(os.PathSeparator)
}

// StringPath returns entire path of the file
func (img *ImagePath) StringPath() string {
	if len(img.Checksum) == 0 {
		return ""
	}

	return img.StringDir() + img.Checksum + img.Extension
}
