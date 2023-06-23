-- +goose Up
-- add column "create_time" to table: "oauth2_client"
ALTER TABLE `oauth2_client` ADD COLUMN `create_time` datetime NOT NULL;
-- add column "update_time" to table: "oauth2_client"
ALTER TABLE `oauth2_client` ADD COLUMN `update_time` datetime NOT NULL;
-- add column "create_time" to table: "shorts"
ALTER TABLE `shorts` ADD COLUMN `create_time` datetime NOT NULL;
-- add column "update_time" to table: "shorts"
ALTER TABLE `shorts` ADD COLUMN `update_time` datetime NOT NULL;
-- add column "create_time" to table: "oauth2_token"
ALTER TABLE `oauth2_token` ADD COLUMN `create_time` datetime NOT NULL;
-- add column "update_time" to table: "oauth2_token"
ALTER TABLE `oauth2_token` ADD COLUMN `update_time` datetime NOT NULL;

-- +goose Down
-- reverse: add column "update_time" to table: "oauth2_token"
ALTER TABLE `oauth2_token` DROP COLUMN `update_time`;
-- reverse: add column "create_time" to table: "oauth2_token"
ALTER TABLE `oauth2_token` DROP COLUMN `create_time`;
-- reverse: add column "update_time" to table: "shorts"
ALTER TABLE `shorts` DROP COLUMN `update_time`;
-- reverse: add column "create_time" to table: "shorts"
ALTER TABLE `shorts` DROP COLUMN `create_time`;
-- reverse: add column "update_time" to table: "oauth2_client"
ALTER TABLE `oauth2_client` DROP COLUMN `update_time`;
-- reverse: add column "create_time" to table: "oauth2_client"
ALTER TABLE `oauth2_client` DROP COLUMN `create_time`;
