package models

type User struct {
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ProjectID int    `json:"project_id"`
	Help      string `json:"help"`
	Message   string `json:"message"`
}
