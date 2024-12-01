package services

import (
	"hackathons-app/internal/models"
	"hackathons-app/internal/repositories"
)

type HackathonService struct {
	repository repositories.HackathonRepository
}

func NewHackathonService(repository repositories.HackathonRepository) HackathonService {
	return HackathonService{repository: repository}
}

func (s *HackathonService) GetAll() ([]models.Hackathon, error) {
	return s.repository.GetAll()
}

func (s *HackathonService) GetAllWithUsers() ([]models.Hackathon, error) {
	return s.repository.GetAllWithUsers()
}

func (s *HackathonService) GetById(id int64) (models.Hackathon, error) {
	return s.repository.GetById(id)
}

func (s *HackathonService) GetWithUsersById(id int64) (models.Hackathon, error) {
	return s.repository.GetWithUsersById(id)
}

func (s *HackathonService) Create(hackathon *models.Hackathon) error {
	return s.repository.Create(hackathon)
}

func (s *HackathonService) CreateMany(hackathons []*models.Hackathon) error {
	return s.repository.CreateMany(hackathons)
}

func (s *HackathonService) Update(hackathon *models.Hackathon) error {
	return s.repository.Update(hackathon)
}

func (s *HackathonService) DeleteById(id int64) error {
	return s.repository.DeleteById(id)
}

func (s *HackathonService) AddUser(hackathonId, userId int64) error {
	return s.repository.AddUser(hackathonId, userId)
}

func (s *HackathonService) RemoveUser(hackathonId, userId int64) error {
	return s.repository.RemoveUser(hackathonId, userId)
}
