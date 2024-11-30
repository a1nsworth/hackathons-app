package repositories

import (
	"hackathons-app/internal/models"
)

// HackathonRepository содержит базовый репозиторий для работы с хакатонами
type HackathonRepository struct {
	baseRepository
}

// GetAllHackathons - получение всех хакатонов
func (r *HackathonRepository) GetAllHackathons() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.db.Find(&hackathons).Error
	return hackathons, err
}

// GetAllHackathonsWithUsers - получение всех хакатонов с пользователями
func (r *HackathonRepository) GetAllHackathonsWithUsers() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := r.db.Preload("Users").Find(&hackathons).Error
	return hackathons, err
}

// GetHackathonById - получение хакатона по ID
func (r *HackathonRepository) GetHackathonById(id int64) (models.Hackathon, error) {
	var hackathon models.Hackathon
	err := r.db.First(&hackathon, id).Error
	return hackathon, err
}

// GetHackathonWithUsersById - получение хакатона с пользователями по ID
func (r *HackathonRepository) GetHackathonWithUsersById(id int64) (models.Hackathon, error) {
	var hackathon models.Hackathon
	err := r.db.Preload("Users").First(&hackathon, id).Error
	return hackathon, err
}

// CreateHackathon - создание хакатона
func (r *HackathonRepository) CreateHackathon(hackathon *models.Hackathon) error {
	return r.db.Create(&hackathon).Error
}

// CreateHackathons - создание нескольких хакатонов
func (r *HackathonRepository) CreateHackathons(hackathons []*models.Hackathon) error {
	return r.db.Create(&hackathons).Error
}

// UpdateHackathon - обновление хакатона
func (r *HackathonRepository) UpdateHackathon(hackathon *models.Hackathon) error {
	return r.db.Save(&hackathon).Error
}

// DeleteHackathonById - удаление хакатона по ID
func (r *HackathonRepository) DeleteHackathonById(id int64) error {
	var hackathon models.Hackathon
	err := r.db.First(&hackathon, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&hackathon).Error
}

// AddUserToHackathon - добавление пользователя в хакатон
func (r *HackathonRepository) AddUserToHackathon(hackathonId int64, userId int64) error {
	var hackathon models.Hackathon
	var user models.User

	if err := r.db.First(&hackathon, hackathonId).Error; err != nil {
		return err
	}
	if err := r.db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := r.db.Model(&hackathon).Association("Users").Append(&user); err != nil {
		return err
	}

	return nil
}

// RemoveUserFromHackathon - удаление пользователя из хакатона
func (r *HackathonRepository) RemoveUserFromHackathon(hackathonId int64, userId int64) error {
	var hackathon models.Hackathon
	var user models.User

	if err := r.db.First(&hackathon, hackathonId).Error; err != nil {
		return err
	}
	if err := r.db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := r.db.Model(&hackathon).Association("Users").Delete(&user); err != nil {
		return err
	}

	return nil
}
