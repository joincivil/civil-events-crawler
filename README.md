![Civil Logo](docs/civil_logo_white.png?raw=true)

---
[Civil](https://joincivil.com/) is a decentralized and censorship resistant ecosystem for online Journalism. Read more in our whitepaper.

This repository contains open-source code to capture and handle Civil-specific smart contract event log data. It is written in `golang`. It currently captures Civil TCR and Civil Newsroom related events, but can be expanded to capture additional events.

[![CircleCI](https://img.shields.io/circleci/project/github/joincivil/civil-events-crawler.svg)](https://circleci.com/gh/joincivil/civil-events-crawler/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/joincivil/civil-events-crawler)](https://goreportcard.com/report/github.com/joincivil/civil-events-crawler)
[![Gitter chat](https://badges.gitter.im/joincivil/Lobby.png)](https://gitter.im/joincivil/Lobby)
[![Telegram chat](https://img.shields.io/badge/chat-telegram-blue.svg)](https://t.me/join_civil)

For Civil's main open-source tools and packages, check out [http://github.com/joincivil/Civil](http://github.com/joincivil/Civil).

## Contributing

Civil's ecosystem is free and open-source, we're all part of it and you're encouraged to be a part of it with us.  We are looking to evolve this into something the community will find helpful and effortless to use.

If you're itching to dwelve deeper inside, [**help wanted**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
and [**good first issue**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) labels are good places to get started and learn the architecture.

## Install Requirements

This project is using `make` to run setup, builds, tests, etc and has been tested and running on `go 1.12.7`.

Ensure that your `$GOPATH` and `$GOROOT` are setup properly in your shell configuration and that this repo is cloned into the appropriate place in the `$GOPATH`. i.e. `$GOPATH/src/github.com/joincivil/civil-events-crawler/`

To setup the necessary requirements:

```
make setup
```

## Code Generation

There are a few places where code/artifacts need to be moved or generated before the project can be built, tested, and/or linted.  This is likely a place that can be streamlined and improved as time goes on.

### Solidity Wrappers

There are Solidity wrappers that are created by `abigen` from the `go-ethereum` package using ABI artifacts in [`joincivil/Civil`](http://github.com/joincivil/Civil).
The wrappers are currently pre-generated and included as a part of the Civil [`joincivil/go-common`](http://github.com/joincivil/go-common) library and imported here as a part of `/vendor` packages.

Please reference the [`joincivil/go-common`](http://github.com/joincivil/go-common) repository for additional information on generating the Civil contract wrappers.

### Contract Types

After any new contract wrappers are imported into the project, to configure event tracking for that contract:

1. Add configuration data to the `ContractTypeToSpecs` map in `model/contracttypes.go`.
2. Add another `ContractType` that represents your new type to the `ContractType` enum in `model/contracttypes.go`.
3. Rebuild the crawler.

In the future, we hope to support a YAML or JSON based type configuration in support of generic reuse of this project.

### Contract Watchers

There is a number of `Watch*` methods for each Civil Solidity contract wrapper that allow us to listen and stream contract events.  The wrappers around these `Watch*` methods are generated using the `cmd/eventhandlergen` command.  These will be placed into the `pkg/generated/watchers` directory.

```
make generate-civil-watchers
```

### Contract Filterers

There is a number of `Filter*` methods for each Civil Solidity contract wrapper that allow us to collect existing contract events.  The wrappers around these `Filter*` methods are generated using the `cmd/eventhandlergen` command.  These will be placed into the `pkg/generated/filterers` directory.

```
make generate-civil-filterers
```

### Event Handler Lists

This creates wrapper functions around each contract's set of filterers and watchers.  A map of contract names to their smart contract address is passed in these functions and determines which set of filterer/watchers need to be started in the crawler.  These are generated using the `cmd/handlerlistgen` command.  These will be placed into the `pkg/generated/handlerlist` directory.

```
make generate-civil-handler-lists
```

### Common

This creates some common code in use for filterers/watchers and other code. These are generated using the `cmd/eventhandlergen` command.  These will be placed into the `pkg/generated/common` directory.

```
make generate-civil-common
```

## Lint

Check all the packages for linting errors using a variety of linters via `golangci-lint`.  Check the `Makefile` for the up to date list of linters.

```
make lint
```

## Build


```
make build
```

## Testing

Runs the tests and checks code coverage across the project. Produces a `coverage.txt` file for use later.

```
make test
```

## Code Coverage Tool

Run `make test` and launches the HTML code coverage tool.

```
make cover
```

## Run

The crawler relies on environment vars for configuration. To configure locally, edit the `.env` file included in the repo to what is needed.

To run the service:

```
go run cmd/crawler/main.go
```

To find all the available configuration environment vars:

```
go run cmd/crawler/main.go -h
```

### Supported Civil Contract Short Names
`civiltcr`, `newsroom`, `civilplcrvoting`, `cvltoken`, `civilparameterizer`, `civilgovernment`, `newsroomfactory`, `multisigfactory`

### Supported Persister Types
`none`, `postgresql`

### Enable Info Logging

Add `-logtostderr=true -stderrthreshold=INFO -v=2` as arguments for the `main.go` command.

## Persistence

The crawler is build to accept an implementation of persistence interfaces as defined in `pkg/model/persisttypes.go`.  These interfaces allow the crawler to store down specific data related to it's operation as well as the events to be collected.

The initial reference implementation will be written for storing the data to `PostgreSQL`. However, the hope to add additional implementations as needed and as the community sees fit.

