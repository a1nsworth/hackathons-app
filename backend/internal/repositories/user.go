package repositories

import (
	"hackathons-app/internal/models"

	"gorm.io/gorm"
)

// UserRepository содержит базовый репозиторий для работы с пользователями
type UserRepository struct {
	baseRepository
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{newBaseRepository(db)}
}

// GetAllUsers - получение всех пользователей
func (r UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// GetAllUsersWithHackathons - получение всех пользователей с хакатонами
func (r UserRepository) GetAllUsersWithHackathons() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Hackathons").Find(&users).Error
	return users, err
}

// GetUserInfoById - получение пользователя по ID
func (r UserRepository) GetUserInfoById(id int64) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

// GetUserHackathonsById - получение хакатонов по ID пользователя
func (r UserRepository) GetUserHackathonsById(id int64) ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.db.Preload("Users").Where("user_id = ?", id).Find(&hackathons).Error
	return hackathons, err
}

// GetUserWithHackathonsById - получение пользователя с хакатонами по ID
func (r UserRepository) GetUserWithHackathonsById(id int64) (models.User, error) {
	var user models.User
	err := r.db.Preload("Hackathons").First(&user, id).Error
	return user, err
}

// CreateUser - создание нового пользователя
func (r UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

// CreateUsers - создание нескольких пользователей
func (r UserRepository) CreateUsers(users []*models.User) error {
	return r.db.Create(&users).Error
}

// UpdateUser - обновление данных пользователя
func (r UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(&user).Error
}

// DeleteUserById - удаление пользователя по ID
func (r UserRepository) DeleteUserById(id int64) error {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}
