-- +goose Up

CREATE TABLE users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;