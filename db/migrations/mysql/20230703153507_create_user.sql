-- +goose Up
-- create "users" table
CREATE TABLE `users` (`id` char(36) NOT NULL, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `subject` varchar(256) NOT NULL, `profile` varchar(255) NOT NULL, `email` varchar(256) NOT NULL, `email_hash` text NOT NULL, `email_verified` bool NOT NULL, `username` varchar(256) NULL, `picture` varchar(255) NULL, `claims` blob NULL, PRIMARY KEY (`id`), UNIQUE INDEX `users_email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- +goose Down
-- reverse: create "users" table
DROP TABLE `users`;
