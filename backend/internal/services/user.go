package services

import (
	"hackathons-app/internal/models"
	"hackathons-app/internal/repositories"
)

type UserService struct {
	userRepo      repositories.UserRepository
	hackathonRepo repositories.HackathonRepository
}

func (s *UserService) UserRepo() repositories.UserRepository {
	return s.userRepo
}

func (s *UserService) HackathonRepo() repositories.HackathonRepository {
	return s.hackathonRepo
}

func NewUserService(
	userRepository repositories.UserRepository,
	hackathonRepository repositories.HackathonRepository,
) UserService {
	return UserService{userRepo: userRepository, hackathonRepo: hackathonRepository}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) GetAllWithHackathons() ([]models.User, error) {
	return s.userRepo.GetAllWithHackathons()
}

func (s *UserService) GetById(id int64) (models.User, error) {
	return s.userRepo.GetById(id)
}

func (s *UserService) GetByEmail(email string) (models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UserService) GetHackathonsById(id int64) ([]models.Hackathon, error) {
	return s.userRepo.GetHackathonsById(id)
}

func (s *UserService) GetWithHackathonsById(id int64) (models.User, error) {
	return s.userRepo.GetWithHackathonsById(id)
}

func (s *UserService) Create(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) CreateMany(users []*models.User) error {
	return s.userRepo.CreateMany(users)
}

func (s *UserService) Update(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteById(id int64) error {
	return s.userRepo.DeleteById(id)
}

func (s *UserService) AddHackathonById(userId, hackathonId int64) (err error) {
	user, err := s.GetById(userId)
	if err != nil {
		return nil
	}
	hackathon, err := s.hackathonRepo.GetById(hackathonId)
	if err != nil {
		return nil
	}

	err = s.hackathonRepo.Connection().GetDB().Model(&user).Association("Hackathons").Append(
		&hackathon,
	)
	if err != nil {
		return err
	}
	return
}
