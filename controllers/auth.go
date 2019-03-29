package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jottsu/sns-sample-api/controllers/params"
	"github.com/jottsu/sns-sample-api/controllers/views"
	"github.com/jottsu/sns-sample-api/models"
	"github.com/jottsu/sns-sample-api/repositories"
	"github.com/jottsu/sns-sample-api/utils"
)

func Signup(c *gin.Context) {
	p := &params.SignupParams{}
	if err := c.BindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Name:  p.Name,
		Email: p.Email,
	}

	tx, err := repositories.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	savedUser, err := repositories.SaveUserWithTx(tx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("faild to save user"))
		return
	}

	userSecret, err := models.CreateUserSecret(savedUser.ID, p.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("faild to create user secret"))
		return
	}

	_, err = repositories.SaveUserSecretWithTx(tx, userSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("faild to save user secret"))
		return
	}

	token, err := utils.CreateUserJwt(savedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("faild to create user jwt"))
		return
	}

	if err := repositories.Commit(tx); err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("faild to commit"))
		return
	}

	c.JSON(http.StatusOK, views.ToSignupView(token, savedUser))
}

func Signin(c *gin.Context) {
	p := &params.SigninParams{}
	if err := c.BindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := repositories.FindUserByEmail(p.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userSecret, err := repositories.FindUserSecretByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := userSecret.CheckPassword(p.Password); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := utils.CreateUserJwt(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, views.ToSignupView(token, user))
}
