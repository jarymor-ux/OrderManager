-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
	id bigserial primary key,
	username varchar not null,
	password_hash varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
