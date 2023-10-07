package repopostgres

import (
	"fmt"
	entityuser "gin/internal/entities/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepoPostgres interface {
	InsertUser(account, password string) error
	QueryUser(account string) ([]entityuser.User, error)
	CloseConn() error
}

type repoPostgres struct {
	DB *gorm.DB
}

func NewRepoPostgres(user, pwd, db, host, port, sslmode string) (RepoPostgres, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, pwd, db, host, port, sslmode)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open error:%s", err.Error())
	}

	return &repoPostgres{
			DB: gormDB,
		},
		nil

}

func (rp *repoPostgres) CloseConn() error {
	sql, err := rp.DB.DB()
	if err != nil {
		return err
	}

	if err = sql.Close(); err != nil {
		return err
	}

	return nil
}
