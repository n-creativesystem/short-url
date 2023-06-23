-- +goose Up
-- modify "oauth2_token" table
ALTER TABLE `oauth2_token` MODIFY COLUMN `code` varchar(512) NULL, MODIFY COLUMN `access` varchar(512) NULL, MODIFY COLUMN `refresh` varchar(512) NULL;

-- +goose Down
-- reverse: modify "oauth2_token" table
ALTER TABLE `oauth2_token` MODIFY COLUMN `refresh` varchar(512) NULL DEFAULT "", MODIFY COLUMN `access` varchar(512) NULL DEFAULT "", MODIFY COLUMN `code` varchar(512) NULL DEFAULT "";
