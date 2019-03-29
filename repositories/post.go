package repositories

import (
	"github.com/jottsu/sns-sample-api/models"
)

func FindPostByID(id uint64) (*models.Post, error) {
	post := &models.Post{}
	if err := DB.First(post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func SavePost(post *models.Post) (*models.Post, error) {
	if err := DB.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
