COMPOSE := COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 CURRENT_UID=$(shell id -u):$(shell id -g) docker-compose

.PHONY: dev
dev: ## Starts a development environment using docker-compose.
	@$(COMPOSE) up -d --remove-orphans
	@$(COMPOSE) logs --tail=100 -t -f $(logs)

.PHONY: destroy
destroy: ## Destroys the docker-compose environment.
	@$(COMPOSE) down -v --remove-orphans
	@$(COMPOSE) rm -vf


.PHONY: build
build: ## Builds all of the images found in the docker-compose configuration.
	@$(COMPOSE) build

test_before: 
	CURRENT_UID=$(shell id -u):$(shell id -g) docker-compose run --rm --publish 8100:8100 saasbackend go test -tags=before -p 1 ./...

test_after: 
	CURRENT_UID=$(shell id -u):$(shell id -g) docker-compose run --rm --publish 8100:8100 saasbackend go test -tags=after -p 1 ./...
