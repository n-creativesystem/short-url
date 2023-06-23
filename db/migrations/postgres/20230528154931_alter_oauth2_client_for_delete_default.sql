-- +goose Up
-- modify "oauth2_token" table
ALTER TABLE "oauth2_token" ALTER COLUMN "code" DROP DEFAULT, ALTER COLUMN "access" DROP DEFAULT, ALTER COLUMN "refresh" DROP DEFAULT;

-- +goose Down
-- reverse: modify "oauth2_token" table
ALTER TABLE "oauth2_token" ALTER COLUMN "refresh" SET DEFAULT '', ALTER COLUMN "access" SET DEFAULT '', ALTER COLUMN "code" SET DEFAULT '';
