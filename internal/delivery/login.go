package handle

import (
	entity_jwt "gin/internal/entities/jwt"
	entityuser "gin/internal/entities/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseData struct {
	Token string `json:"token"`
}

// @Summary  register account
// @Tags Account
// @Description  註冊帳號
// @Produce application/json
// @Accept application/json
// @Router /user [post]
// @Failure 400 object ResponseError
// @Success 200 object ResponseData
// @Param  user body  entityuser.User true "register account"
func (a *ApiHandler) RegisterAccount(c *gin.Context) {

	usr := &entityuser.User{}
	if err := c.ShouldBindJSON(usr); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	log.Printf("user:%+v \n", usr)

	if err := a.Usecase.RegisterAccount(strings.TrimSpace(usr.Account), strings.TrimSpace(usr.Password)); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	newtoken := entity_jwt.NewJwtAnalysis()
	usrId, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	t, err := newtoken.CreateToken(usrId, usr.Account, time.Now().Add(time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseData{t})

}

// @Summary login
// @Tags Account
// @Description 帳密登入認證
// @Produce json
// @Accept  json
// @Param  login body entityuser.User true "login"
// @Router /login [post]
// @Failure  400 object  ResponseError
// @Success  200 object  ResponseData
func (a *ApiHandler) Login(c *gin.Context) {
	usr := &entityuser.User{}
	err := c.ShouldBindJSON(usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ok, err := a.Usecase.Login(usr.Account, usr.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if !ok {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	newtoken := entity_jwt.NewJwtAnalysis()
	usrId, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	t, err := newtoken.CreateToken(usrId, usr.Account, time.Now().Add(time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseData{t})

}
