SCHEMA=""
CREATE_SCHEMA_CMD_EXAMPLE="make gen/ent-schema SCHEMA=ToDo"
.PHONY: gen/ent-schema
gen/ent-schema:
	@if test -z $${SCHEMA}; then echo "$(CREATE_SCHEMA_CMD_EXAMPLE)" && exit 1; fi
	go run -mod=mod entgo.io/ent/cmd/ent new --target pkg/infrastructure/rdb/ent/schema $(SCHEMA)

MFILE=""
DRIVER?="mysql" # postgres(sqlboiler: psql), mysql
CREATE_MIGRATION_CMD_EXAMPLE='make migration-file MFILE=something_to_do'
## create migration sql file
.PHONY: migration-file
migration-file:
	@if test -z $${MFILE}; then echo -e "$(CREATE_MIGRATION_CMD_EXAMPLE)" && exit 1; fi
	go run main.go migrator create $(MFILE) --driver $(DRIVER)

.PHONY: migration-file/mysql
migration-file/mysql:
	make migration-file DRIVER=mysql

.PHONY: migration-file/postgres
migration-file/postgres:
	make migration-file DRIVER=postgres

.PHONY: migration-file/sqlite
migration-file/sqlite:
	make migration-file DRIVER=sqlite

.PHONY: gen/migrate
gen/migrate: migration-file/mysql migration-file/postgres migration-file/sqlite


CREATE_MIGRATION_CMD_EXAMPLE_BY_DYNAMODB='make migration-file/dynamodb MFILE=something_to_do'
.PHONY: migration-file/dynamodb
migration-file/dynamodb:
	@if test -z $${MFILE}; then echo -e "$(CREATE_MIGRATION_CMD_EXAMPLE_BY_DYNAMODB)" && exit 1; fi
	FILE="./db/migrations/dynamodb/$$(date "+%Y%m%d%H%M%S")_$${MFILE}.sh"; \
	echo "#!/usr/bin/env sh" > "$${FILE}"; \
	chmod +x "$${FILE}"

.PHONY: migration-up
migration-up:
	go run main.go migrator up --driver $(DRIVER)

.PHONY: migration-up/mysql
migration-up/mysql:
	make migration-up DRIVER=mysql

.PHONY: migration-up/postgres
migration-up/postgres:
	make migration-up DRIVER=postgres

.PHONY: migration-up/sqlite
migration-up/sqlite:
	make migration-up DRIVER=sqlite

.PHONY: migration-up/all
migration-up/all: migration-up/mysql migration-up/postgres migration-up/sqlite

.PHONY: migration-down
migration-down:
	go run main.go migrator down --driver $(DRIVER)

.PHONY: generate-jwt-key
generate-jwt-key:
	openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out pkg/api/handler/files/jwt-private.pem
	openssl pkey -in pkg/api/handler/files/jwt-private.pem -pubout > pkg/api/handler/files/jwt-public.pem

CREATE_MIGRATION_CMD_EXAMPLE_BY_TEST_DATA='make migration-file/testdata MFILE=something_to_do'
## create migration sql file
.PHONY: migration-file/testdata
migration-file/testdata:
	@if test -z $${MFILE}; then echo -e "$(CREATE_MIGRATION_CMD_EXAMPLE_BY_TEST_DATA)" && exit 1; fi
	goose -dir ./db/testdata create $(MFILE) sql

.PHONY: migration-up/testdata
migration-up/testdata:
	go run main.go migrator up --dir db/testdata --driver $(DRIVER)

.PHONY: migration-down/testdata
migration-down/testdata:
	go run main.go migrator down --dir db/testdata --driver $(DRIVER)

.PHONY: test
test:
	export $$(cat .env .test.env | grep -v ^#) && go test ./...

.PHONY: test/coverage
test/coverage:
	export $$(cat .env .test.env | grep -v ^#) && go test  ./... -coverpkg=./... -cover -coverprofile=cover/cover.out.tmp
	cat cover/cover.out.tmp | grep -v "/mock/" | grep -v "/infrastructure/rdb/ent/" | grep -v "/tests/" > cover/cover.out
	rm -f cover/cover.out.tmp
	go tool cover -html=cover/cover.out -o cover/cover.html
	go tool cover -func=cover/cover.out -o cover/cover.txt

generate/mock:
	mockgen -package=oauth2 -destination=pkg/mock/external/oauth2/token_store.go github.com/go-oauth2/oauth2/v4 TokenStore

.PHONY: gql/gen/backend
gql/gen/backend:
	go run -mod=mod github.com/99designs/gqlgen generate

.PHONY: gql/gen/frontend
gql/gen/frontend:
	cd frontend && yarn gql:gen

.PHONY: gql/gen
gql/gen: gql/gen/backend gql/gen/frontend

MODE=""
INFO=""
TAGS=""
.PHONY: swag/gen
swag/gen:
	go run -mod=mod github.com/swaggo/swag/v2/cmd/swag init \
		-o docs/openapi/$(MODE) \
		--generalInfo $(INFO) \
		--tags $(TAGS) \
		--pd \
		-v3.1

.PHONY: swag/gen/api
swag/gen/api:
	make swag/gen \
		MODE=api \
		INFO=pkg/interfaces/router/gin_api_router.go \
		TAGS=API

.PHONY: swag/gen/ui
swag/gen/ui:
	make swag/gen \
		MODE=ui \
		INFO=pkg/interfaces/router/gin_ui_router.go \
		TAGS=UI
	cd frontend && yarn swag:gen

# .PHONY: swag/gen/main
# swag/gen/main:
# 	make swag/gen MODE=main

RUN_DRIVER?="mysql"
RUN_PORT?=3000
.PHONY: run/uiserver
run/uiserver:
	export $$(cat .env | grep -v ^#) && go run main.go server web-ui --port $(RUN_PORT) --driver $(RUN_DRIVER) --mode debug

.PHONY: run/apiserver
run/apiserver:
	export $$(cat .env | grep -v ^#) && go run main.go server api --port $(RUN_PORT) --driver $(RUN_DRIVER) --mode debug
