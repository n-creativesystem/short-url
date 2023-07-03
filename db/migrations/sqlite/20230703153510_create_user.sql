-- +goose Up
-- create "users" table
CREATE TABLE `users` (`id` uuid NOT NULL, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `subject` text NOT NULL, `profile` text NOT NULL, `email` text NOT NULL, `email_hash` text NOT NULL, `email_verified` bool NOT NULL, `username` text NULL, `picture` text NULL, `claims` blob NULL, PRIMARY KEY (`id`));
-- create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);

-- +goose Down
-- reverse: create index "users_email" to table: "users"
DROP INDEX `users_email`;
-- reverse: create "users" table
DROP TABLE `users`;
