BUILD_DIR = ./bin
SRC_DIR = ./cmd
NAME_SERVER = friends-server
NAME_CLI = friends-cli
NAME_LB = friends-balancer

build = cli

.PHONY: all
all: clean up

.PHONY: up
up: $(build) down
	docker compose up -d

.PHONY: cli
cli: vendor
	@echo $@ build stage
	CGO_ENABLED=0 GO111MODULE=on go build -o $(BUILD_DIR)/$(NAME_CLI) $(SRC_DIR)/$(NAME_CLI)/*.go

.PHONY: vendor
vendor:
	@echo vendoring...
	go mod tidy
	go mod vendor

.PHONY: down
down:
	docker compose down

.PHONY: clean
clean:
	@echo erase bin files
	rm -f $(BUILD_DIR)/$(NAME_CLI)

	@echo erase database
	rm -rf database/*

	#@echo erase docker containers
	#@docker stop $(NAME_MONGO) || echo "no containers running"
	#@docker rm $(NAME_MONGO) || echo "no containers found"

	#@echo erase images
