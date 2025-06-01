-- +goose Up
-- +goose StatementBegin
CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    discription TEXT NOT NULL,
    photo VARCHAR(256) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd