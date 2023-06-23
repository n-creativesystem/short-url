-- +goose Up
-- modify "oauth2_client" table
ALTER TABLE `oauth2_client` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
-- modify "oauth2_token" table
ALTER TABLE `oauth2_token` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
-- modify "shorts" table
ALTER TABLE `shorts` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;

-- +goose Down
-- reverse: modify "shorts" table
ALTER TABLE `shorts` DROP COLUMN `update_time`, DROP COLUMN `create_time`;
-- reverse: modify "oauth2_token" table
ALTER TABLE `oauth2_token` DROP COLUMN `update_time`, DROP COLUMN `create_time`;
-- reverse: modify "oauth2_client" table
ALTER TABLE `oauth2_client` DROP COLUMN `update_time`, DROP COLUMN `create_time`;
