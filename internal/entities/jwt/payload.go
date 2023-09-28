package entity_jwt

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	ExpiredTime time.Time `json:"expired_time"`
}

func (p *Payload) Valid() error { //實現interface 這邊是要通過token的payload來判斷是否是合理的token

	if time.Now().After(p.ExpiredTime) { //查看token是否過期
		return fmt.Errorf("this token has expired.")
	}

	return nil

}
