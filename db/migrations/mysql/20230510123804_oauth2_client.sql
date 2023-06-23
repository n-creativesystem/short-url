-- +goose Up
-- create "oauth2_client" table
CREATE TABLE `oauth2_client` (`id` varchar(255) NOT NULL, `secret` varchar(255) NOT NULL, `domain` varchar(255) NOT NULL, `public` bool NOT NULL, `user_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- +goose Down
-- reverse: create "oauth2_client" table
DROP TABLE `oauth2_client`;
