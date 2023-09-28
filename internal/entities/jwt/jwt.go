package entity_jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtAnalysis struct {
	secretKey string
}

func NewJwtAnalysis() JwtAnalysis {

	return JwtAnalysis{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (j *JwtAnalysis) CreateToken(userId uuid.UUID, name string, expiredTime time.Time) (string, error) {

	payload := &Payload{
		UserID:      userId,
		Name:        name,
		ExpiredTime: expiredTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //選擇簽名演算法(加密方式)跟輸入自訂的payload
	signedToken, err := token.SignedString([]byte(j.secretKey)) //最後一步用secretKey簽署token,會轉[]byte是因為go的secrectKey預期是[]byte類型
	if err != nil {
		return "", fmt.Errorf("SignedString error:%s", err.Error())
	}
	return signedToken, nil

}

func (j *JwtAnalysis) VerifyToken(token string) (*Payload, error) {

	payload := &Payload{}

	keyFunc := func(t *jwt.Token) (interface{}, error) { //該func 主要功能 回傳secretKey, 會轉[]byte是因為go的secrectKey預期是[]byte類型
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { //非必要但有更好,提早過濾非約定加密方式,通過interface斷言來判斷該token的加密方式是不是我們使用的方法,上方使用SigningMethodHS256算法加密
			return nil, fmt.Errorf("sign method type error")
		}
		return []byte(j.secretKey), nil

	}

	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc) //開始解析token是否正確
	if err != nil {
		return nil, fmt.Errorf("token parse with claim error:%s", err.Error())
	}

	payLoadInfo, ok := jwtToken.Claims.(*Payload) //讀取token的payload
	if !ok {
		return nil, fmt.Errorf("claim payload failed")
	}
	return payLoadInfo, nil
}

func CheckTokenFormat(token string) (string, error) {
	if len(strings.TrimSpace(token)) == 0 {
		return "", fmt.Errorf("token is empty")
	}

	tokenField := strings.Fields(strings.TrimSpace(token))

	if len(tokenField) != 2 {
		return "", fmt.Errorf("token length error")
	}

	if strings.ToLower(tokenField[0]) != "bearer" {
		return "", fmt.Errorf("this token is not bear type")
	}

	return tokenField[1], nil

}
