package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/jottsu/sns-sample-api/models"
)

func SaveUserSecretWithTx(tx *gorm.DB, userSecret *models.UserSecret) (*models.UserSecret, error) {
	if err := tx.Create(userSecret).Error; err != nil {
		return nil, err
	}
	return userSecret, nil
}

func FindUserSecretByUserID(userID uint64) (*models.UserSecret, error) {
	userSecret := &models.UserSecret{}
	if err := DB.Where(&models.UserSecret{UserID: userID}).First(&userSecret).Error; err != nil {
		return nil, err
	}
	return userSecret, nil
}
