package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"golang-server/database"
	"golang-server/model"
	"io"
	"os"
	"strconv"
	"strings"
)

var CompetitorService = &competitorService{repository: &database.MysqlCompetitorRepository{}}

type competitorService struct {
	repository *database.MysqlCompetitorRepository
}

func (svc *competitorService) FindAll(competitors *[]model.Competitor, condition uint64) (*[]model.Competitor, error) {
	match := new(model.Match)
	match.ID = condition
	err := svc.repository.FindAll(competitors, match)
	return competitors, err
}

func (svc *competitorService) FindOne(id uint64) (*model.Competitor, error) {
	competitor := new(model.Competitor)
	competitor.ID = id
	err := svc.repository.FindOne(&competitor)
	return competitor, err
}

func (svc *competitorService) Save(competitors *[]model.Competitor) (*[]model.Competitor, error) {
	err := svc.repository.SaveAll(*competitors)
	return competitors, err
}

func (svc *competitorService) SaveFile(src io.Reader, ext string) (path string, err error) {
	var buffer bytes.Buffer
	hasher := sha256.New()
	if _, err = io.Copy(hasher, io.TeeReader(src, &buffer)); err != nil {
		return
	}

	hash := hasher.Sum(nil)
	shard := strconv.FormatUint(binary.BigEndian.Uint64(hash)%997, 10)
	dir := strings.Join([]string{"image", shard}, string(os.PathSeparator))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.FileMode(0775))
	}

	checksum := hex.EncodeToString(hash)
	path = dir + string(os.PathSeparator) + checksum + ext
	if _, err = os.Stat(path); os.IsExist(err) {
		return
	}

	var saved *os.File
	saved, err = os.Create(path)
	defer saved.Close()
	if err != nil {
		return
	}

	if _, err = io.Copy(saved, &buffer); err != nil {
		return
	}
	return
}
