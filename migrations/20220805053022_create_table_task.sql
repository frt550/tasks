-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.task (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    is_completed boolean DEFAULT false NOT NULL,
    created_at timestamp NOT NULL,
    completed_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.task;
-- +goose StatementEnd
