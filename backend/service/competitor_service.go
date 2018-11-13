package service

import (
	"bytes"
	"errors"
	"golang-server/database"
	"golang-server/model"
	"io"
)

var CompetitorService = &competitorService{repository: &database.MysqlCompetitorRepository{}}

type competitorService struct {
	repository *database.MysqlCompetitorRepository
}

func (svc *competitorService) FindLatest(competitors *[]model.Competitor, criteria *model.Competitor) (*[]model.Competitor, error) {
	return competitors, svc.repository.FindWithCursor(competitors, criteria)
}

func (svc *competitorService) Save(competitor *model.Competitor, match *model.Match) (*model.Competitor, error) {
	var count int
	if err := svc.repository.Count(&count, &model.Competitor{MatchID: competitor.MatchID}); err != nil {
		return nil, err
	}
	if count >= match.Quota {
		return nil, errors.New("더 이상 추가하실 수 없습니다")
	}
	return competitor, svc.repository.Save(competitor)
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

func (svc *competitorService) Update(updated *model.Competitor) error {
	// TODO: user > match > competitor
	original := new(model.Competitor)
	original.ID = updated.ID
	if err := svc.repository.FindOne(original); err != nil {
		return err
	}
	return svc.repository.Update(updated, &model.Competitor{Caption: updated.Caption})
}
