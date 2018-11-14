package service

import "golang-server/database/mysql"

var Playing = &playingService{mysql: mysql.BaseMysqlRepository}

type playingService struct {
	mysql *mysql.MysqlRepository
}

func (svc *playingService) generateGame() error {
	return nil
}
