# mongodb-adapter

## Table of Contents

- [mongodb-adapter](#mongodb-adapter)
  - [Table of Contents](#table-of-contents)
  - [About ](#about-)
  - [Getting Started ](#getting-started-)
    - [Prerequisites](#prerequisites)
    - [Installing](#installing)
      - [Go download and install](#go-download-and-install)
      - [Podman download and install](#podman-download-and-install)
      - [Make install](#make-install)

## About <a name = "about"></a>

MongoDB adapter to migrate data from MongoDB to Google Cloud Spanner. While requring the structure/schema of mongosb database.

## Getting Started <a name = "getting_started"></a>

- Run `go mod tidy` to install neccesery libraries required.
- Run MAKE commands in the sequential matter as provided in the MAKEFILE.
- Instruction while executing `gorun command`:
      - `make gorun MONGODB_ADDRESS=<YOUR MONGO DATABASE ADDRESS LINK> MONGODB_DATABASE=<YOUR MONGO DATABASE NAME>`


### Prerequisites

- Golang 1.15 or above
- Docker/podman
- Make
### Installing
#### Go download and install
    [Golang Installation](https://golang.org/dl/)
#### Podman download and install
    For MAC, if HOMEREW is installed:
        `brew install podman`
    For other OS refer:
        [Podman installation](https://podman.io/docs/installation)
#### Make install
    `brew install make`