// Package esdcadminmodule provides admin functionality for the ESDC backend
package esdcadminmodule

import (
	"github.com/aruncs31s/esdcadminmodule/handler"
	"github.com/aruncs31s/esdcadminmodule/routes"
	"github.com/aruncs31s/esdcadminmodule/service"
)

// Re-export main types
type AdminHandler = handler.AdminHandler
type AdminService = service.AdminService

// NewAdminHandler creates a new admin handler
var NewAdminHandler = handler.NewAdminHandler

// NewAdminService creates a new admin service
var NewAdminService = service.NewAdminService

// RegisterAdminRoutes registers admin routes
var RegisterAdminRoutes = routes.RegisterAdminRoutes
