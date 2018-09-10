POSTGRES_DATA_DIR=postgresdata
POSTGRES_DOCKER_IMAGE=circleci/postgres:9.6-alpine
POSTGRES_PORT=5432
POSTGRES_DB_NAME=civil_crawler
POSTGRES_USER=docker
POSTGRES_PSWD=docker

GOCMD=go
GOGEN=$(GOCMD) generate
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVER=$(GOCMD) tool cover
ABIGEN=abigen

## Check to see if these commands are installed
GO:=$(shell command -v go 2> /dev/null)
DOCKER:=$(shell command -v docker 2> /dev/null)
APT:=$(shell command -v apt-get 2> /dev/null)

ABI_DIR=abi

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_CONTRACT_DIR=$(GENERATED_DIR)/contract
GENERATED_WATCHER_DIR=$(GENERATED_DIR)/watcher
GENERATED_FILTERER_DIR=$(GENERATED_DIR)/filterer
GENERATED_COMMON_DIR=$(GENERATED_DIR)/common
GENERATED_HANDLER_LIST_DIR=$(GENERATED_DIR)/handlerlist

## Civil specific commands
EVENTHANDLER_GEN_MAIN=cmd/eventhandlergen/main.go
HANDLERLIST_GEN_MAIN=cmd/handlerlistgen/main.go

## Reliant on go and $GOPATH being set.
.PHONY: check-go-env
check-go-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif

## NOTE: If installing on a Mac, use Docker for Mac, not Docker toolkit
## https://www.docker.com/docker-mac
.PHONY: check-docker-env
check-docker-env:
ifndef DOCKER
	$(error docker command is not installed or in PATH)
endif

.PHONY: install-dep
install-dep: check-go-env ## Installs dep
	@mkdir -p $(GOPATH)/bin
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

.PHONY: install-linter
install-linter: check-go-env ## Installs linter
	@$(GOGET) -u github.com/alecthomas/gometalinter
	@gometalinter --install
ifdef APT
	@sudo apt-get install golang-race-detector-runtime || true
endif

.PHONY: install-cover
install-cover: check-go-env ## Installs code coverage tool
	@$(GOGET) -u golang.org/x/tools/cmd/cover

.PHONY: install-abigen
install-abigen: check-go-env ## Installs the Ethereum abigen tool
	@$(GOGET) -u github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: setup
setup: check-go-env install-dep install-linter install-cover install-abigen ## Sets up the tooling.

.PHONY: postgres-setup-launch
postgres-setup-launch:
ifeq ("$(wildcard $(POSTGRES_DATA_DIR))", "")
	mkdir -p $(POSTGRES_DATA_DIR)
	docker run \
		-v $$PWD/$(POSTGRES_DATA_DIR):/tmp/$(POSTGRES_DATA_DIR) -i -t $(POSTGRES_DOCKER_IMAGE) \
		/bin/bash -c "cp -rp /var/lib/postgresql /tmp/$(POSTGRES_DATA_DIR)"
endif
	docker run -e "POSTGRES_USER="$(POSTGRES_USER) -e "POSTGRES_PASSWORD"=$(POSTGRES_PSWD) -e "POSTGRES_DB"=$(POSTGRES_DB_NAME) \
	    -v $$PWD/$(POSTGRES_DATA_DIR)/postgresql:/var/lib/postgresql -d -p $(POSTGRES_PORT):$(POSTGRES_PORT) \
		$(POSTGRES_DOCKER_IMAGE);

.PHONY: postgres-check-available
postgres-check-available:
	@for i in `seq 1 10`; \
	do \
		nc -z localhost 5432 2> /dev/null && exit 0; \
		sleep 3; \
	done; \
	exit 1;

.PHONY: postgres-start
postgres-start: check-docker-env postgres-setup-launch postgres-check-available ## Starts up a development PostgreSQL server
	@echo "Postgresql launched and available"

.PHONY: postgres-stop
postgres-stop: check-docker-env ## Stops the development PostgreSQL server
	@docker stop `docker ps -q`
	@echo 'Postgres stopped'

## gometalinter config in .gometalinter.json
.PHONY: lint
lint: ## Runs linting.
	@gometalinter ./...

.PHONY: generate-civil
generate: generate-contracts generate-civil-watchers generate-civil-filterers generate-civil-handler-lists ## Runs all the civil code generation

.PHONY: generate-civil-watchers
generate-civil-watchers: ## Runs watchergen to generate contract Watch* wrapper code for Civil.
	@mkdir -p $(GENERATED_WATCHER_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr watcher watcher > ./$(GENERATED_WATCHER_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom watcher watcher > ./$(GENERATED_WATCHER_DIR)/newsroom.go

.PHONY: generate-civil-filterers
generate-civil-filterers: ## Runs filterergen to generate contract Filter* wrapper code for Civil.
	@mkdir -p $(GENERATED_FILTERER_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr filterer filterer > ./$(GENERATED_FILTERER_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom filterer filterer > ./$(GENERATED_FILTERER_DIR)/newsroom.go

.PHONY: generate-civil-common
generate-civil-common: ## Runs commongen to generate common contract wrapper code for Civil.
	@mkdir -p $(GENERATED_COMMON_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr common common > ./$(GENERATED_COMMON_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom common common > ./$(GENERATED_COMMON_DIR)/newsroom.go

.PHONY: generate-civil-handler-lists
generate-civil-handler-lists: ## Runs handlerlistgen to generate handler list wrapper code for Civil.
	@mkdir -p $(GENERATED_HANDLER_LIST_DIR)
	@$(GORUN) $(HANDLERLIST_GEN_MAIN) handlerlist > ./$(GENERATED_HANDLER_LIST_DIR)/handlerlist.go

.PHONY: generate-civil-contracts
generate-civil-contracts: ## Builds the contract wrapper code from the ABIs in /abi for Civil.
ifneq ("$(wildcard $(ABI_DIR)/*.abi)", "")
	@mkdir -p $(GENERATED_CONTRACT_DIR)
	@$(ABIGEN) -abi ./$(ABI_DIR)/CivilTCR.abi -bin ./$(ABI_DIR)/CivilTCR.bin -type CivilTCRContract -out ./$(GENERATED_CONTRACT_DIR)/CivilTCRContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Newsroom.abi -bin ./$(ABI_DIR)/Newsroom.bin -type NewsroomContract -out ./$(GENERATED_CONTRACT_DIR)/NewsroomContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/PLCRVoting.abi -bin ./$(ABI_DIR)/PLCRVoting.bin -type PLCRVotingContract -out ./$(GENERATED_CONTRACT_DIR)/PLCRVotingContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Parameterizer.abi -bin ./$(ABI_DIR)/Parameterizer.bin -type ParameterizerContract -out ./$(GENERATED_CONTRACT_DIR)/ParameterizerContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Government.abi -bin ./$(ABI_DIR)/Government.bin -type GovernmentContract -out ./$(GENERATED_CONTRACT_DIR)/GovernmentContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/EIP20.abi -bin ./$(ABI_DIR)/EIP20.bin -type EIP20Contract -out ./$(GENERATED_CONTRACT_DIR)/EIP20.go -pkg contract
else
	$(error No abi files found; copy them to /abi after generation)
endif

.PHONY: build
build: ## Builds the code.
	$(GOBUILD) -o ./build/crawler cmd/crawler/main.go

.PHONY: test
test: ## Runs unit tests and tests code coverage.
	@echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: test-integration
test-integration: ## Runs tagged integration tests
	@echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=60s -tags=integration ./...

.PHONY: cover
cover: test ## Runs unit tests, code coverage, and runs HTML coverage tool.
	@$(GOCOVER) -html=coverage.txt

.PHONY: clean
clean: ## go clean and clean up of artifacts.
	@$(GOCLEAN) ./... || true
	@rm coverage.txt || true

## Some magic from http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
