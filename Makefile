NAME := app
DEV_CONFIG_PATH := ./configs/dev.yml
CONFIG_TEMPLATE_PATH := ./configs/template.yml

# Docker
DOCKER_APP_FILENAME ?= deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml

# sed
SECRET_KEY ?= "very-secret-key"
CONFIG_PATH ?= ./configs/new.yml

# Build
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd/${NAME}

define sedi
    sed --version >/dev/null 2>&1 && sed -- $(1) > ${CONFIG_PATH} || sed "" $(1) > ${CONFIG_PATH}
endef

.PHONY: run
run: entgen
	go run cmd/$(NAME)/main.go ${DEV_CONFIG_PATH}

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes

.PHONY: entgen
entgen:
	cd internal/repository/entgo && go generate ./ent

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: up
up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up --build

.PHONY: lint
lint:
	echo lint

.PHONY: test
test:
	go test ./...

.PHONY: substitute_config_vars
substitute_config_vars:
	$(call sedi," \
		s/{{db_password}}/${DB_PASSWORD}/g; \
		s/{{db_name}}/${DB_NAME}/g;         \
		s/{{db_host}}/${DB_HOST}/g;         \
		s/{{db_port}}/${DB_PORT}/g;         \
		s/{{db_user}}/${DB_USER}/g;         \
		s/{{secret_key}}/${SECRET_KEY}/g;   \
		" ${CONFIG_TEMPLATE_PATH})
	echo ${CONFIG_PATH}
	cat ${CONFIG_PATH}