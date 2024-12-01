package services

import (
	"hackathons-app/internal/models"
	"hackathons-app/internal/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserService{repository: repository}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) GetAllWithHackathons() ([]models.User, error) {
	return s.repository.GetAllWithHackathons()
}

func (s *UserService) GetById(id int64) (models.User, error) {
	return s.repository.GetById(id)
}

func (s *UserService) GetHackathonsById(id int64) ([]models.Hackathon, error) {
	return s.repository.GetHackathonsById(id)
}

func (s *UserService) GetWithHackathonsById(id int64) (models.User, error) {
	return s.repository.GetWithHackathonsById(id)
}

func (s *UserService) Create(user *models.User) error {
	return s.repository.Create(user)
}

func (s *UserService) CreateMany(users []*models.User) error {
	return s.repository.CreateMany(users)
}

func (s *UserService) Update(user *models.User) error {
	return s.repository.Update(user)
}

func (s *UserService) DeleteById(id int64) error {
	return s.repository.DeleteById(id)
}
