-- +goose Up
-- +goose StatementBegin
CREATE TABLE administrators (
    id SERIAL PRIMARY KEY,  -- Уникальный идентификатор администратора
    name VARCHAR(255) NOT NULL,  -- Имя администратора
    role VARCHAR(100) NOT NULL   -- Роль пользователя
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS administrators;
-- +goose StatementEnd