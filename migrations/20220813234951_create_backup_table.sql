-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.backup (
    id serial PRIMARY KEY,
    data text NOT NULL,
    created_at timestamp(0) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.backup;
-- +goose StatementEnd
