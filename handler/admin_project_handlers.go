package handler

import (
	"strconv"

	"github.com/aruncs31s/esdcadminmodule/service"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

// ProjectServiceInterface defines what we need from project service
type ProjectServiceInterface interface {
	CreateProject(user string, project ProjectCreationDTO) (interface{}, error)
}

// ProjectCreationDTO represents project creation request
type ProjectCreationDTO struct {
	Title        string    `json:"title" example:"My Project"`
	Image        *string   `json:"image" example:"https://example.com/image.jpg"`
	Description  string    `json:"description" example:"This is a sample project description"`
	GithubLink   string    `json:"github_link" example:"https://github.com/user/project"`
	Tags         *[]string `json:"tags" example:"go,api,backend"`
	Contributers *[]string `json:"contributers" example:"user1,user2,user3"`
	Technologies *[]string `json:"technologies" example:"Go, Gin, GORM"`
	LiveUrl      *string   `json:"live_url" example:"https://example.com/live"`
	Category     string    `json:"category" example:"Web Development"`
}

type AdminProjectHandlers interface {
	GetAllProjects(c *gin.Context)
	GetProjectByID(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type adminProjectHandler struct {
	adminService   service.AdminService
	projectService ProjectServiceInterface
	responseHelper responsehelper.ResponseHelper
}

func newAdminProjectHandler(projectService ProjectServiceInterface, adminService service.AdminService) AdminProjectHandlers {
	responseHelper := responsehelper.NewResponseHelper()
	return &adminProjectHandler{
		responseHelper: responseHelper,
		projectService: projectService,
		adminService:   adminService,
	}
}

func verifyAdminRole(c *gin.Context, responseHelper responsehelper.ResponseHelper) bool {
	role := c.GetString("role")
	if role != "admin" {
		responseHelper.Unauthorized(c, "Admin role required. Your role: "+role)
		return false
	}
	return true
}
func getPaginationParams(c *gin.Context) (string, string) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "50")
	return pageStr, limitStr
}
func getPaginationParamsInt(c *gin.Context) (int, int) {
	page, limit := getPaginationParams(c)
	pageInt := 1
	limitInt := 50
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageInt = p
	}
	if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 1000 {
		limitInt = l
	}
	return limitInt, (pageInt - 1) * limitInt
}

func (h *adminProjectHandler) GetAllProjects(c *gin.Context) {
	if !verifyAdminRole(c, h.responseHelper) {
		return
	}
	limit, offset := getPaginationParamsInt(c)

	allProjects, err := h.adminService.GetProjectsEssentialInfo(limit, offset)
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve projects", err)
		return
	}
	h.responseHelper.Success(c, allProjects)
}
func (h *adminProjectHandler) GetProjectByID(c *gin.Context) {
	// Implementation here
}
func (h *adminProjectHandler) CreateProject(c *gin.Context) {
	// Implementation here
	if !verifyAdminRole(c, h.responseHelper) {
		return
	}
	userName := c.GetString("user")
	if userName == "" {
		h.responseHelper.Unauthorized(c, "User not authenticated")
		return
	}
	var projectDTO ProjectCreationDTO
	if err := c.ShouldBindJSON(&projectDTO); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	createdProject, err := h.projectService.CreateProject(userName, projectDTO)

	if err != nil || createdProject == nil { // Error while creating project
		h.responseHelper.InternalError(c, "Failed to create project", err)
		return
	}
	h.responseHelper.Success(c, createdProject)

}
func (h *adminProjectHandler) UpdateProject(c *gin.Context) {
	// Implementation here
}
func (h *adminProjectHandler) DeleteProject(c *gin.Context) {
	// Implementation here
}
