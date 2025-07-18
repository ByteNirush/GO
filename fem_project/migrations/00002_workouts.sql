-- +goose Up
-- +goose Statement
CREATE TABLE IF NOT EXISTS workouts(
    id BIGSERIAL PRIMARY KEY,
    -- user_id
    title VARCHAR(100) NOT NULL,
    description TEXT,
    duration_minutes INT NOT NULL,
    calories_burned INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Statement

-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd