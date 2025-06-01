-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    project_id INT NOT NULL,
    help TEXT,
    message TEXT,
    CONSTRAINT fk_project
        FOREIGN KEY(project_id) 
        REFERENCES projects(project_id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd