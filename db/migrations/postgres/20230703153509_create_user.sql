-- +goose Up
-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "subject" character varying NOT NULL, "profile" character varying NOT NULL, "email" character varying NOT NULL, "email_hash" text NOT NULL, "email_verified" boolean NOT NULL, "username" character varying NULL, "picture" character varying NULL, "claims" bytea NULL, PRIMARY KEY ("id"));
-- create index "users_email" to table: "users"
CREATE UNIQUE INDEX "users_email" ON "users" ("email");

-- +goose Down
-- reverse: create index "users_email" to table: "users"
DROP INDEX "users_email";
-- reverse: create "users" table
DROP TABLE "users";
