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

func (svc *competitorService) Save(competitor *model.Competitor) (*model.Competitor, error) {
	err := svc.repository.Save(competitor)
	competitor.ID = competitor.BaseModel.ID
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
