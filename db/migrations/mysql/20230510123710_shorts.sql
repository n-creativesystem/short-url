-- +goose Up
-- create "shorts" table
CREATE TABLE `shorts` (`id` bigint NOT NULL AUTO_INCREMENT, `key` varchar(255) NOT NULL, `url` varchar(255) NOT NULL, `author` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- +goose Down
-- reverse: create "shorts" table
DROP TABLE `shorts`;
