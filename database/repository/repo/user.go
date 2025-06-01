package repo

import (
	"server/database/models"
	"context"


)

func (r *Repository) CreateUser(ctx context.Context, u models.User) (models.User, error) {
    const q = `
        INSERT INTO users (name, email, project_id, help, message)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING user_id, name, email, project_id, help, message
    `

    var created models.User
    err := r.Client.QueryRow(
        ctx, q,
        u.Name,
        u.Email,
        u.ProjectID,
        u.Help,
        u.Message,
    ).Scan(
        &created.UserID,
        &created.Name,
        &created.Email,
        &created.ProjectID,
        &created.Help,
        &created.Message,
    )
    if err != nil {
        return models.User{}, err
    }

    return created, nil
}