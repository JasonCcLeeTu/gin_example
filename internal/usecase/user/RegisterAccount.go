package usecaseusr

import (
	"fmt"
	"gin/internal/pkg/encryption"
	repopostgres "gin/internal/repository/db/postgres"
	"strings"

	"github.com/pkg/errors"
)

type UsecaseUser interface {
	RegisterAccount(account, pwd string) error
	Login(account, pwd string) (bool, error)
}

func NewUsecaseUser(myPostgres repopostgres.RepoPostgres) UsecaseUser {
	return &usecaseUser{
		RepoPostgres: myPostgres,
	}
}

type usecaseUser struct {
	RepoPostgres repopostgres.RepoPostgres
}

func (u *usecaseUser) RegisterAccount(account, pwd string) error {

	if err := u.RepoPostgres.InsertUser(account, pwd); err != nil {
		return errors.Wrap(err, "InsertUser error")
	}
	return nil
}

func (u *usecaseUser) Login(account, pwd string) (bool, error) {
	usr, err := u.RepoPostgres.QueryUser(account)
	if err != nil {
		return false, errors.Wrap(err, "postgres repository  queryUser error")
	}
	if len(usr) == 0 {
		return false, fmt.Errorf("this account is not found")
	}

	err = encryption.ComparePwd(usr[0].Password, strings.TrimSpace(pwd))
	if err != nil {
		return false, errors.Wrap(err, "compare pwd error")
	}

	return true, nil
}
