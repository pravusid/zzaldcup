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

func (svc *competitorService) Save(competitor *model.Competitor, match *model.Match) error {
	var count int
	if err := svc.repository.Count(&count, &model.Competitor{MatchID: competitor.MatchID}); err != nil {
		return err
	}
	if count >= match.Quota {
		return errors.New("error: sufficient competitors")
	}
	return svc.repository.Save(competitor)
}

func (svc *competitorService) SaveFile(src io.Reader, ext string) (*model.ImagePath, error) {
	var path *model.ImagePath

	var buffer bytes.Buffer
	hash, err := FileService.HashingAndBuffering(&src, &buffer)
	if err != nil {
		return path, errors.New("error: file has fault")
	}

	var existence bool
	if path, existence = FileService.GenerateFilePath(hash, "image", ext); existence {
		return path, nil
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
