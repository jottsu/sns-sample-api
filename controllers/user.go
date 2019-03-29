package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jottsu/sns-sample-api/controllers/params"
	"github.com/jottsu/sns-sample-api/repositories"
)

func UserShow(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := repositories.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UserUpdate(c *gin.Context) {
	p := &params.UserUpdateParams{}
	if err := c.BindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := repositories.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("failed to find user"))
		return
	}

	user.Name = p.Name
	user.Email = p.Email
	updatedUser, err := repositories.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("failed to update user"))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
