# Simple Analytic

Simple backend for text based analytic/log created with TDD and clean architecture in mind.

## Overview

* Support only text based log
* Support count and group log based on time

## Project structure

    .
    ├── delivery    # contains all user facing code (rest, GraphQL, GRPC)
    ├── entity      # business logic and entity
    ├── interfaces  # all API and struct interface definition
    ├── pkg         # helper
    ├── repository  # persisted data source (database, redis)
    ├── service     # contains app logic
    ├── .env        # configuration
    ├── Makefile    # task runner
    └── main.go     # entry point

## Getting started

Make sure you can run `Makefile` with `make`.

* To install tools dependency: `make tool`
* To run project: `make run`
* To build project: `make build`
* To generate interface mocks: `make mock`
* To run test: `make test`
* To run static analysis check: `make lint`