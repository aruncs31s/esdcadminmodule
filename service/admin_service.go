package service

import (
	"time"

	"github.com/aruncs31s/esdcadminmodule/dto"
	"github.com/aruncs31s/esdcadminmodule/service/mapper"
)

// UsersStats represents user statistics
type UsersStats struct {
	TotalUsers      int `json:"total_users"`
	TotalProjects   int `json:"total_projects"`
	TotalChallenges int `json:"total_challenges"`
	ActiveUsers     int `json:"active_users"`
}

// ProjectsEssentialInfo represents essential project info for admin panel
type ProjectsEssentialInfo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedBy string `json:"created_by"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserRepositoryInterface defines what we need from user repository
type UserRepositoryInterface interface {
	GetAllUsers() (interface{}, error)
	GetUsersCount() (int, error)
	DeleteUserByID(id uint) error
	CreateUser(user interface{}) error
}

// ProjectRepositoryInterface defines what we need from project repository
type ProjectRepositoryInterface interface {
	GetEssentialInfo(limit, offset int) ([]ProjectsEssentialInfo, error)
	GetProjectsCount() (int, error)
}

type AdminService interface {
	GetProjectsEssentialInfo(limit, offset int) ([]ProjectsEssentialInfo, error)
	GetAllUsers() (*[]dto.UserDataForAdmin, error)
	GetUsersStats() (*UsersStats, error)
	DeleteUser(userID int) error
	CreateUser(user dto.AdminRegisterRequest) error
}

type adminService struct {
	userRepo    UserRepositoryInterface
	projectRepo ProjectRepositoryInterface
}

func NewAdminService(userRepo UserRepositoryInterface, projectRepo ProjectRepositoryInterface) AdminService {
	return &adminService{
		userRepo:    userRepo,
		projectRepo: projectRepo,
	}
}

func (s *adminService) GetProjectsEssentialInfo(limit, offset int) ([]ProjectsEssentialInfo, error) {
	projects, err := s.projectRepo.GetEssentialInfo(limit, offset)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *adminService) GetAllUsers() (*[]dto.UserDataForAdmin, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	filteredUsers := mapper.MapToUserDataForAdmin(users)
	return filteredUsers, nil
}

func (s *adminService) GetUsersStats() (*UsersStats, error) {
	usersCount, err := s.userRepo.GetUsersCount()
	if err != nil {
		return nil, err
	}
	projectsCount, err := s.projectRepo.GetProjectsCount()
	if err != nil {
		return nil, err
	}
	var activeUser = 2
	var totalChallenges = 5

	data := getUserStats(usersCount, projectsCount, totalChallenges, activeUser)

	return &data, nil
}

func getUserStats(usersCount int, projectsCount int, totalChallenges int, activeUser int) UsersStats {
	data := UsersStats{
		TotalUsers:      usersCount,
		TotalProjects:   projectsCount,
		TotalChallenges: totalChallenges,
		ActiveUsers:     activeUser,
	}
	return data
}

func (s *adminService) DeleteUser(userID int) error {
	err := s.userRepo.DeleteUserByID(uint(userID))
	if err != nil {
		return err
	}
	return nil
}

func (s *adminService) CreateUser(user dto.AdminRegisterRequest) error {
	newUser := getUserData(user)
	err := s.userRepo.CreateUser(&newUser)
	if err != nil {
		return err
	}
	return nil
}

// UserData represents user data for creation
type UserData struct {
	Name      string
	Username  string
	Email     string
	Role      string
	Password  string
	Github    *GithubData
	CreatedAt int64
	UpdatedAt int64
}

// GithubData represents github user data
type GithubData struct {
	Username string
}

func getUserData(user dto.AdminRegisterRequest) UserData {
	newUser := UserData{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
		Github: &GithubData{
			Username: user.GithubUsername,
		},
		CreatedAt: time.Time{}.Unix(),
		UpdatedAt: time.Time{}.Unix(),
	}
	return newUser
}
