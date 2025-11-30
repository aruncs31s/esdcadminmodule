package mapper

import (
	"time"

	"github.com/aruncs31s/esdcadminmodule/dto"
)

// User represents the user model for mapping
type User struct {
	ID        uint
	Name      string
	Email     string
	Username  string
	Github    *Github
	Role      string
	Status    string
	CreatedAt int64
	UpdatedAt int64
}

// Github represents github info
type Github struct {
	Username string
}

func MapToUserDataForAdmin(users interface{}) *[]dto.UserDataForAdmin {
	if users == nil {
		return nil
	}

	// Try to assert to slice of User
	userSlice, ok := users.(*[]User)
	if !ok {
		return nil
	}

	var filteredUsers []dto.UserDataForAdmin
	for _, user := range *userSlice {
		githubUsername := ""
		if user.Github != nil {
			githubUsername = user.Github.Username
		}
		filteredUsers = append(filteredUsers, dto.UserDataForAdmin{
			ID:             user.ID,
			Name:           user.Name,
			Email:          user.Email,
			Username:       user.Username,
			GithubUsername: githubUsername,
			Role:           user.Role,
			Status:         user.Status,
			CreatedAt:      getCreatedDateFromNumber(user.CreatedAt),
			UpdatedAt:      getCreatedDateFromNumber(user.UpdatedAt),
		})
	}
	return &filteredUsers
}

func getCreatedDateFromNumber(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
