package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jottsu/sns-sample-api/controllers/params"
	"github.com/jottsu/sns-sample-api/models"
	"github.com/jottsu/sns-sample-api/repositories"
)

func PostCreate(c *gin.Context) {
	p := &params.PostCreateParams{}
	if err := c.BindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	post := &models.Post{
		Text: p.Text,
	}
	savedPost, err := repositories.SavePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, savedPost)
}

func PostShow(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := repositories.FindPostByID(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, post)
}
