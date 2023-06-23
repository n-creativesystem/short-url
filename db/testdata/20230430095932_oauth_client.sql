-- +goose Up
-- +goose StatementBegin
INSERT INTO `oauth2_client` (`id`, `secret`, `domain`, `public`, `user_id`) values ('example', 'example_secret', 'http://localhost:8888', 0, 'user1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `oauth2_client` where `id` = 'example';
-- +goose StatementEnd
