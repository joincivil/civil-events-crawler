![Civil Logo](docs/civil_logo_white.png?raw=true)

---
[Civil](https://joincivil.com/) is a decentralized and censorship resistant ecosystem for online Journalism. Read more in our whitepaper.

This repository contains open-source code to capture and handle Civil-specific smart contract event log data. It is written in `golang`. It currently captures Civil TCR and Civil Newsroom related events, but can be expanded to capture additional events.

For Civil's main open-source tools and packages, check out [http://github.com/joincivil/Civil](http://github.com/joincivil/Civil).

## Contributing

Civil's ecosystem is free and open-source, we're all part of it and you're encouraged to be a part of it with us.  We are looking to evolve this into something the community will find helpful and effortless to use.

If you're itching to dwelve deeper inside, [**help wanted**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
and [**good first issue**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) labels are good places to get started and learn the architecture.

## Install Requirements

This project is using `make` to run setup, builds, tests, etc.  

Ensure that your `$GOPATH` and `$GOROOT` are setup properly in your shell configuration and that this repo is cloned into the appropriate place in the `$GOPATH`. i.e. `$GOPATH/src/github.com/joincivil/civil-events-crawler/`

To setup the necessary requirements:

```
# vgo
# gometalinter
# go tool cover
# abigen

make setup
```

### Dependencies

Uses `vgo` for dependency management, although we may keep a `/vendor/` and use `go get` for backwards compatibility with `go` and some tooling may still rely on `$GOPATH`. **Need to solidify what we are doing here.**


## Lint

Check all the packages for linting errors using a variety of linters via `gometalinter`.  Check the `Makefile` for the up to date list of linters.

```
make lint
```

## Build


```
make build
```

## Testing

Runs the tests across the project.

```
make test
```

## Code Coverage

Run the unit tests and code coverage tool.

```
make cover
```


