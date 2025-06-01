package dto

import "server/database/models"

type CreateUserRequest struct {
	Name      string `form:"name"`
	Email     string `form:"email"`
	ProjectID int    `form:"project_id"`
	Help      string `form:"help"`
	Message   string `form:"message"`
}

func (r CreateUserRequest) ToModel() models.User {
	return models.User{
		Name:      r.Name,
		Email:     r.Email,
		ProjectID: r.ProjectID,
		Help:      r.Help,
		Message:   r.Message,
	}
}
