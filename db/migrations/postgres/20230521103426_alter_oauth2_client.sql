-- +goose Up
-- modify "oauth2_client" table
ALTER TABLE "oauth2_client" ADD COLUMN "app_name" character varying NOT NULL;

-- +goose Down
-- reverse: modify "oauth2_client" table
ALTER TABLE "oauth2_client" DROP COLUMN "app_name";
