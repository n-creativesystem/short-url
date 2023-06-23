-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_oauth2_token" table
CREATE TABLE `new_oauth2_token` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `expired_at` integer NOT NULL, `code` text NULL, `access` text NULL, `refresh` text NULL, `data` text NULL);
-- copy rows from old table "oauth2_token" to new temporary table "new_oauth2_token"
INSERT INTO `new_oauth2_token` (`id`, `expired_at`, `code`, `access`, `refresh`, `data`) SELECT `id`, `expired_at`, `code`, `access`, `refresh`, `data` FROM `oauth2_token`;
-- drop "oauth2_token" table after copying rows
DROP TABLE `oauth2_token`;
-- rename temporary table "new_oauth2_token" to "oauth2_token"
ALTER TABLE `new_oauth2_token` RENAME TO `oauth2_token`;
-- create index "idx_code" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_code` ON `oauth2_token` (`code`);
-- create index "idx_access" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_access` ON `oauth2_token` (`access`);
-- create index "idx_refresh" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_refresh` ON `oauth2_token` (`refresh`);
-- create index "idx_expired_at" to table: "oauth2_token"
CREATE UNIQUE INDEX `idx_expired_at` ON `oauth2_token` (`expired_at`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "idx_expired_at" to table: "oauth2_token"
DROP INDEX `idx_expired_at`;
-- reverse: create index "idx_refresh" to table: "oauth2_token"
DROP INDEX `idx_refresh`;
-- reverse: create index "idx_access" to table: "oauth2_token"
DROP INDEX `idx_access`;
-- reverse: create index "idx_code" to table: "oauth2_token"
DROP INDEX `idx_code`;
-- reverse: create "new_oauth2_token" table
DROP TABLE `new_oauth2_token`;
