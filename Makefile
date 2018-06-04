GOCMD=vgo
GOGEN=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVER=$(GOCMD) tool cover

## Check to see if `go` is installed
GO := $(shell command -v go 2> /dev/null)

ABI_DIR=abi

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_TCR_DIR=$(GENERATED_DIR)/tcr
GENERATED_NEWSROOM_DIR=$(GENERATED_DIR)/newsroom
GENERATED_PLCR_DIR=$(GENERATED_DIR)/plcr
GENERATED_PARAM_DIR=$(GENERATED_DIR)/parameterizer
GENERATED_GOVT_DIR=$(GENERATED_DIR)/government
GENERATED_TOKEN_DIR=$(GENERATED_DIR)/eip20
GENERATED_DIRS=$(GENERATED_TCR_DIR) $(GENERATED_NEWSROOM_DIR) $(GENERATED_PLCR_DIR) \
			   $(GENERATED_GOVT_DIR) $(GENERATED_PARAM_DIR) $(GENERATED_TOKEN_DIR)

## Reliant on go and $GOPATH being set.
check-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif

.PHONY: install-vgo
install-vgo: check-env ## Installs vgo
	go get -u golang.org/x/vgo

.PHONY: install-linter
install-linter: check-env ## Installs linter
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: install-cover
install-cover: check-env ## Installs code coverage tool
	go get -u golang.org/x/tools/cmd/cover

.PHONY: install-abigen
install-abigen: check-env ## Installs the Ethereum abigen tool
	go get -u github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: setup
setup: check-env install-vgo install-linter install-cover install-abigen ## Sets up the golang environment

.PHONY: lint
lint: generate ## Runs linting
	gometalinter \
		--disable-all \
		--enable=golint \
		--enable=gofmt \
		--enable=gotype \
		--enable=vet \
		--enable=deadcode \
		--enable=megacheck \
		--enable=varcheck \
		--enable=structcheck \
		--enable=unconvert \
		--skip=generated \
		--skip=go \
		--deadline=3m \
		--concurrency=2 \
		./...

.PHONY: generate
generate: ## Runs 'go generate' to produce any code that needs generation
	$(GOGEN) ./...

.PHONY: generate-contracts
generate-contracts: ## Builds the contract wrapper code from the ABIs in /abi
ifneq ("$(wildcard $(ABI_DIR)/*.abi)", "")
	mkdir -p $(GENERATED_DIRS)
	abigen -abi ./$(ABI_DIR)/CivilTCR.abi -bin ./$(ABI_DIR)/CivilTCR.bin -type CivilTCRContract -out ./$(GENERATED_TCR_DIR)/CivilTCRContract.go -pkg tcr
	abigen -abi ./$(ABI_DIR)/Newsroom.abi -bin ./$(ABI_DIR)/Newsroom.bin -type NewsroomContract -out ./$(GENERATED_NEWSROOM_DIR)/NewsroomContract.go -pkg newsroom
	abigen -abi ./$(ABI_DIR)/PLCRVoting.abi -bin ./$(ABI_DIR)/PLCRVoting.bin -type PLCRVotingContract -out ./$(GENERATED_PLCR_DIR)/PLCRVotingContract.go -pkg plcr
	abigen -abi ./$(ABI_DIR)/Parameterizer.abi -bin ./$(ABI_DIR)/Parameterizer.bin -type ParameterizerContract -out ./$(GENERATED_PARAM_DIR)/ParameterizerContract.go -pkg parameterizer
	abigen -abi ./$(ABI_DIR)/Government.abi -bin ./$(ABI_DIR)/Government.bin -type GovernmentContract -out ./$(GENERATED_GOVT_DIR)/GovernmentContract.go -pkg government
	abigen -abi ./$(ABI_DIR)/EIP20.abi -bin ./$(ABI_DIR)/EIP20.bin -type EIP20Contract -out ./$(GENERATED_TOKEN_DIR)/EIP20.go -pkg eip20
else
	$(error No abi files found; copy them to /abi after generation)
endif

.PHONY: build
build: generate ## Builds the code
	$(GOBUILD) ./...

.PHONY: test
test: generate ## Runs unit tests
	echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: generate test ## Runs unit tests and checks code coverage
	$(GOCOVER) -html=coverage.txt

.PHONY: clean
clean: ## go clean and clean up of artifacts
	rm -rf pkg/generated
	$(GOCLEAN)

## Some magic from http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
