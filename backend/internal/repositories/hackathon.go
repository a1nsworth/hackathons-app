package repositories

import (
	"hackathons-app/internal/db"
	"hackathons-app/internal/models"
)

// HackathonRepository содержит базовый репозиторий для работы с хакатонами
type HackathonRepository struct {
	baseRepository
}

func NewHackathonRepository(gormDb *db.GormDatabase) HackathonRepository {
	return HackathonRepository{newBaseRepository(gormDb)}
}

// GetAll - получение всех хакатонов
func (r *HackathonRepository) GetAll() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.connection.GetDB().Find(&hackathons).Error
	return hackathons, err
}

// GetAllWithUsers - получение всех хакатонов с пользователями
func (r *HackathonRepository) GetAllWithUsers() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.connection.GetDB().Preload("Users").Find(&hackathons).Error
	return hackathons, err
}

// GetById - получение хакатона по ID
func (r *HackathonRepository) GetById(id int64) (models.Hackathon, error) {
	var hackathon models.Hackathon
	err := r.connection.GetDB().First(&hackathon, id).Error
	return hackathon, err
}

// GetWithUsersById - получение хакатона с пользователями по ID
func (r *HackathonRepository) GetWithUsersById(id int64) (models.Hackathon, error) {
	var hackathon models.Hackathon
	err := r.connection.GetDB().Preload("Users").First(&hackathon, id).Error
	return hackathon, err
}

// Create - создание хакатона
func (r *HackathonRepository) Create(hackathon *models.Hackathon) error {
	return r.connection.GetDB().Create(&hackathon).Error
}

// CreateMany - создание нескольких хакатонов
func (r *HackathonRepository) CreateMany(hackathons []*models.Hackathon) error {
	return r.connection.GetDB().Create(&hackathons).Error
}

// Update - обновление хакатона
func (r *HackathonRepository) Update(hackathon *models.Hackathon) error {
	return r.connection.GetDB().Save(&hackathon).Error
}

// DeleteById - удаление хакатона по ID
func (r *HackathonRepository) DeleteById(id int64) error {
	var hackathon models.Hackathon
	err := r.connection.GetDB().First(&hackathon, id).Error
	if err != nil {
		return err
	}
	return r.connection.GetDB().Delete(&hackathon).Error
}

// AddUser - добавление пользователя в хакатон
func (r *HackathonRepository) AddUser(hackathonId int64, userId int64) error {
	var hackathon models.Hackathon
	var user models.User

	if err := r.connection.GetDB().First(&hackathon, hackathonId).Error; err != nil {
		return err
	}
	if err := r.connection.GetDB().First(&user, userId).Error; err != nil {
		return err
	}

	if err := r.connection.GetDB().Model(&hackathon).Association("Users").Append(&user); err != nil {
		return err
	}

	return nil
}

// RemoveUser - удаление пользователя из хакатона
func (r *HackathonRepository) RemoveUser(hackathonId int64, userId int64) error {
	var hackathon models.Hackathon
	var user models.User

	if err := r.connection.GetDB().First(&hackathon, hackathonId).Error; err != nil {
		return err
	}
	if err := r.connection.GetDB().First(&user, userId).Error; err != nil {
		return err
	}

	if err := r.connection.GetDB().Model(&hackathon).Association("Users").Delete(&user); err != nil {
		return err
	}

	return nil
}
