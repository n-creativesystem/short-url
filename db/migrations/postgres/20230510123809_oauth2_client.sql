-- +goose Up
-- create "oauth2_client" table
CREATE TABLE "oauth2_client" ("id" character varying NOT NULL, "secret" character varying NOT NULL, "domain" character varying NOT NULL, "public" boolean NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"));

-- +goose Down
-- reverse: create "oauth2_client" table
DROP TABLE "oauth2_client";
