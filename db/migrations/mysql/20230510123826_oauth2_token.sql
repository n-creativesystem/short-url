-- +goose Up
-- create "oauth2_token" table
CREATE TABLE `oauth2_token` (`id` bigint NOT NULL AUTO_INCREMENT, `expired_at` bigint NOT NULL, `code` varchar(512) NULL DEFAULT "", `access` varchar(512) NULL DEFAULT "", `refresh` varchar(512) NULL DEFAULT "", `data` text NULL, PRIMARY KEY (`id`), UNIQUE INDEX `idx_access` (`access`), UNIQUE INDEX `idx_code` (`code`), UNIQUE INDEX `idx_expired_at` (`expired_at`), UNIQUE INDEX `idx_refresh` (`refresh`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- +goose Down
-- reverse: create "oauth2_token" table
DROP TABLE `oauth2_token`;
