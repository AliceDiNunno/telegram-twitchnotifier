package postgres

import (
	"TwitchNotifierForTelegram/src/domain"
	"gorm.io/gorm"
)

type User struct {
	ID          int64
	Username    string
	Displayname string
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userRepo {
	return userRepo{
		db: db,
	}
}

func (r userRepo) RegisterUser(userID int64, username string, displayName string) error {
	result := r.db.Create(&User{
		ID:          userID,
		Username:    username,
		Displayname: displayName,
	})

	return result.Error
}

func (r userRepo) UpdateUser(userID int64, username string, displayName string) error {
	result := r.db.Model(&User{}).Where("id = ?", userID).Updates(User{
		Username:    username,
		Displayname: displayName,
	})

	return result.Error
}

func (r userRepo) GetUser(userID int64) (domain.User, error) {
	var user User

	result := r.db.First(&user, "id = ?", userID)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return domain.User{
		ID:          user.ID,
		Username:    user.Username,
		Displayname: user.Displayname,
	}, nil
}
