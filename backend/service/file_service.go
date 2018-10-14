package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
)

var FileService = &fileService{}

type fileService struct{}

func (fileService) HashingAndBuffering(source *io.Reader, buffer *bytes.Buffer) ([]byte, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, io.TeeReader(*source, buffer)); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func (fileService) GenerateFilePath(hash []byte, baseDir string, extension string) (string, bool) {
	shard := strconv.FormatUint(binary.BigEndian.Uint64(hash)%997, 10)
	dir := strings.Join([]string{baseDir, shard}, string(os.PathSeparator))
	if baseDir == "" {
		dir = strings.TrimPrefix(dir, string(os.PathSeparator))
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.FileMode(0775))
	}

	checksum := hex.EncodeToString(hash)
	path := dir + string(os.PathSeparator) + checksum + extension
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path, false
	}
	return path, true
}

func (fileService) CreateFile(path string, data io.Reader) error {
	target, err := os.Create(path)
	defer target.Close()
	if err != nil {
		return err
	}

	if _, err = io.Copy(target, data); err != nil {
		return err
	}
	return nil
}
