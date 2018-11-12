package service

import (
	"bytes"
	"golang-server/database"
	"golang-server/model"
	"io"
)

var CompetitorService = &competitorService{repository: &database.MysqlCompetitorRepository{}}

type competitorService struct {
	repository *database.MysqlCompetitorRepository
}

func (svc *competitorService) FindLatest(competitors *[]model.Competitor, criteria *model.Competitor) (*[]model.Competitor, error) {
	err := svc.repository.FindWithCursor(competitors, criteria)
	return competitors, err
}

func (svc *competitorService) Save(competitor *model.Competitor) (*model.Competitor, error) {
	err := svc.repository.Save(competitor)
	return competitor, err
}

func (svc *competitorService) SaveFile(src io.Reader, ext string) (path *model.ImagePath, err error) {
	var buffer bytes.Buffer
	hash, err := FileService.HashingAndBuffering(&src, &buffer)

	var existence bool
	if path, existence = FileService.GenerateFilePath(hash, "image", ext); existence {
		return
	}

	return path, FileService.CreateFile(path, &buffer)
}
