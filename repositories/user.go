package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/jottsu/sns-sample-api/models"
)

func FindUserByID(id uint64) (*models.User, error) {
	user := &models.User{}
	if err := DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := DB.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUser(user *models.User) (*models.User, error) {
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUserWithTx(tx *gorm.DB, user *models.User) (*models.User, error) {
	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *models.User) (*models.User, error) {
	if err := DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
