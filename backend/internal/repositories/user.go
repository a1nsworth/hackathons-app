package repositories

import (
	"hackathons-app/internal/db"
	"hackathons-app/internal/models"
)

// UserRepository содержит базовый репозиторий для работы с пользователями
type UserRepository struct {
	baseRepository
}

func NewUserRepository(gormDb *db.GormDatabase) UserRepository {
	return UserRepository{newBaseRepository(gormDb)}
}

// GetAll - получение всех пользователей
func (r UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.connection.GetDB().Find(&users).Error
	return users, err
}

// GetAllWithHackathons - получение всех пользователей с хакатонами
func (r UserRepository) GetAllWithHackathons() ([]models.User, error) {
	var users []models.User
	err := r.connection.GetDB().Preload("Hackathons").Find(&users).Error
	return users, err
}

// GetById - получение пользователя по ID
func (r UserRepository) GetById(id int64) (models.User, error) {
	var user models.User
	err := r.connection.GetDB().First(&user, id).Error
	return user, err
}

// GetHackathonsById - получение хакатонов по ID пользователя
func (r UserRepository) GetHackathonsById(id int64) ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.connection.GetDB().Preload("Users").Where("user_id = ?", id).Find(&hackathons).Error
	return hackathons, err
}

// GetWithHackathonsById - получение пользователя с хакатонами по ID
func (r UserRepository) GetWithHackathonsById(id int64) (models.User, error) {
	var user models.User
	err := r.connection.GetDB().Preload("Hackathons").First(&user, id).Error
	return user, err
}

// Create - создание нового пользователя
func (r UserRepository) Create(user *models.User) error {
	return r.connection.GetDB().Create(&user).Error
}

// CreateMany - создание нескольких пользователей
func (r UserRepository) CreateMany(users []*models.User) error {
	return r.connection.GetDB().Create(&users).Error
}

// Update - обновление данных пользователя
func (r UserRepository) Update(user *models.User) error {
	return r.connection.GetDB().Save(&user).Error
}

// DeleteById - удаление пользователя по ID
func (r UserRepository) DeleteById(id int64) error {
	var user models.User
	err := r.connection.GetDB().First(&user, id).Error
	if err != nil {
		return err
	}
	return r.connection.GetDB().Delete(&user).Error
}

func (r UserRepository) GetByEmail(email string) (user models.User, err error) {
	err = r.connection.GetDB().Where("email = ?", email).First(&user).Error
	return
}
