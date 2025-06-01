package repo

import (
	"context"
	"server/database/models"
)

func (r *Repository) GetProjects(ctx context.Context) ([]models.Project, error) {
	q := `
    SELECT project_id, name, discription, photo FROM Projets;
  `

	rows, err := r.Client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Projects []models.Project

	for rows.Next() {
		var wh models.Project
		err := rows.Scan(
			&wh.ProjectID,
			&wh.Name,
			&wh.Discription,
			&wh.Photo,
		)

		if err != nil {
			return nil, err
		}
		Projects = append(Projects, wh)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Projects, nil
}
