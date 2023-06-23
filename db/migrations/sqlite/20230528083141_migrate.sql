-- +goose Up
-- create "oauth2_client" table
CREATE TABLE `oauth2_client` (`id` text NOT NULL, `secret` text NOT NULL, `domain` text NOT NULL, `public` bool NOT NULL, `user_id` text NOT NULL, `app_name` text NOT NULL, PRIMARY KEY (`id`));
-- create "oauth2_token" table
CREATE TABLE `oauth2_token` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `expired_at` integer NOT NULL, `code` text NULL DEFAULT '', `access` text NULL DEFAULT '', `refresh` text NULL DEFAULT '', `data` text NULL);
-- create index "idx_code" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_code` ON `oauth2_token` (`code`);
-- create index "idx_access" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_access` ON `oauth2_token` (`access`);
-- create index "idx_refresh" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_refresh` ON `oauth2_token` (`refresh`);
-- create index "idx_expired_at" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_expired_at` ON `oauth2_token` (`expired_at`);
-- create "shorts" table
CREATE TABLE `shorts` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `key` text NOT NULL, `url` text NOT NULL, `author` text NOT NULL);

-- +goose Down
-- reverse: create "shorts" table
DROP TABLE `shorts`;
-- reverse: create index "idx_expired_at" to table: "oauth2_token"
DROP INDEX `idx_expired_at`;
-- reverse: create index "idx_refresh" to table: "oauth2_token"
DROP INDEX `idx_refresh`;
-- reverse: create index "idx_access" to table: "oauth2_token"
DROP INDEX `idx_access`;
-- reverse: create index "idx_code" to table: "oauth2_token"
DROP INDEX `idx_code`;
-- reverse: create "oauth2_token" table
DROP TABLE `oauth2_token`;
-- reverse: create "oauth2_client" table
DROP TABLE `oauth2_client`;
