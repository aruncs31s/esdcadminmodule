package dto

// AdminRegisterRequest represents admin user registration request
// @Description Admin user registration request payload
type AdminRegisterRequest struct {
	Name           string `json:"name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	GithubUsername string `json:"github_username"`
	Role           string `json:"role" example:"admin"` // User role (e.g., "admin", "user")
}

type UserDataForAdmin struct {
	ID             uint   `json:"id"`              // User ID
	Name           string `json:"name"`            // Full name
	Email          string `json:"email"`           // Email address
	Username       string `json:"username"`        // Username
	GithubUsername string `json:"github_username"` // GitHub username
	Role           string `json:"role"`            // User role (e.g., "admin", "user")
	Status         string `json:"status"`          // Account status (active/inactive)
	CreatedAt      string `json:"created_at"`      // Account creation timestamp
	UpdatedAt      string `json:"updated_at"`      // Account last update timestamp
}
