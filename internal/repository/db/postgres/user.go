package repopostgres

import (
	"fmt"
	entityuser "gin/internal/entities/user"
	"gin/internal/pkg/encryption"
	"log"
)

func (rp *repoPostgres) InsertUser(account, password string) error {

	hashedPwd, err := encryption.Encrypt(password, 12)
	if err != nil {
		return fmt.Errorf("Encrypt error:%s", err.Error())
	}

	usr := &entityuser.User{
		Account:  account,
		Password: string(hashedPwd),
	}
	log.Printf("User: %+v \n", usr)
	if err := rp.DB.Table("user").Create(usr).Error; err != nil {
		return fmt.Errorf("repository insert data error:%s", err.Error())
	}
	return nil
}

func (rp *repoPostgres) QueryUser(account string) ([]entityuser.User, error) {
	var usr []entityuser.User
	if err := rp.DB.Table("user").Where("account=?", account).Find(&usr).Error; err != nil {
		return nil, fmt.Errorf("query data error:%s", err.Error())
	}

	return usr, nil

}
