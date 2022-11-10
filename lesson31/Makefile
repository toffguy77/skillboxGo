NAME_SERVER = friends-server
NAME_CLI = friends-cli
NAME_LB = friends-balancer
NAME_MONGO = friends-mongo
BUILD_DIR = ./bin
SRC_DIR = ./cmd
MONGO_ADMIN = thatguy
MONGO_PASS = password

fullBuild := server cli balancer

.PHONY: all
all: clean mongo startServer startLB

.PHONY: server
server: vendor
	@echo $@ build stage
	CGO_ENABLED=0 GO111MODULE=on go build -o $(BUILD_DIR)/$(NAME_SERVER) $(SRC_DIR)/$(NAME_SERVER)/*.go

.PHONY: cli
cli: vendor
	@echo $@ build stage
	CGO_ENABLED=0 GO111MODULE=on go build -o $(BUILD_DIR)/$(NAME_CLI) $(SRC_DIR)/$(NAME_CLI)/*.go

.PHONY: balancer
balancer: vendor
	@echo $@ build stage
	CGO_ENABLED=0 GO111MODULE=on go build -o $(BUILD_DIR)/$(NAME_LB) $(SRC_DIR)/$(NAME_LB)/*.go

.PHONY: vendor
vendor:
	@echo vendoring...
	go mod tidy
	go mod vendor

.PHONY: mongo
mongo:
	docker compose up -d

.PHONY: startServer
startServer: $(fullBuild)
	@echo starting 2 instances of the server
	./bin/$(NAME_SERVER) &

.PHONY: startLB
startLB: $(fullBuild)
	@echo starting load balancer
	./bin/$(NAME_LB) &

.PHONY: clean
clean:
	@echo erase bin files
	pgrep -o $(NAME_SERVER) | xargs kill -9 || echo "no running $(NAME_SERVER)"
	rm -f $(BUILD_DIR)/$(NAME_SERVER)
	pgrep -o $(NAME_LB) | xargs kill -9 || echo "no running $(NAME_LB)"
	rm -f $(BUILD_DIR)/$(NAME_LB)
	rm -f $(BUILD_DIR)/$(NAME_CLI)
	@echo erase database
	rm -rf ./database/*
	@echo erase docker container
	docker compose down
	@docker stop $(NAME_MONGO) || echo "no containers running"
	@docker rm $(NAME_MONGO) || echo "no containers found"

