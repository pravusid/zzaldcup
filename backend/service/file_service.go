package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"golang-server/model"
	"io"
	"os"
	"strconv"
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

func (fileService) GenerateFilePath(hash []byte, baseDir string, extension string) (*model.ImagePath, bool) {
	path := model.ImagePath{
		BaseDir:   baseDir,
		Shard:     strconv.FormatUint(binary.BigEndian.Uint64(hash)%997, 10),
		Checksum:  hex.EncodeToString(hash),
		Extension: extension,
	}

	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		os.Mkdir(baseDir, os.FileMode(0775))
	}

	dir, _ := path.StringDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.FileMode(0775))
	}

	strPath, _ := path.StringPath()
	if _, err := os.Stat(strPath); os.IsNotExist(err) {
		return &path, false
	}
	return &path, true
}

func (fileService) CreateFile(path *model.ImagePath, data io.Reader) error {
	strPath, pathErr := path.StringPath()
	if pathErr != nil {
		return pathErr
	}

	target, err := os.Create(strPath)
	defer target.Close()
	if err != nil {
		return err
	}

	if _, err = io.Copy(target, data); err != nil {
		return err
	}
	return nil
}
