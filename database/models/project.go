package models

type Project struct {
	ProjectID   int64  `json:"project_id" `
	Name        string `json:"name"`
	Discription string `json:"discription"`
	Photo       string `json:"photo"`
}
