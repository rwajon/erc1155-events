# ERC1155-events

ERC1155-events is an application that subscribes to smart contract transfer events and store them.

# Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

# Prerequisites

- [Go(1.17)](https://go.dev/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

# Install and Run the Project

```bash
# clone the repo
$ git clone https://github.com/rwajon/erc1155-events.git

$ cd erc1155-events

# install go modules
$ go mod download

# run the development server
$ ./scripts/run.sh
```

### using docker & docker compose

```bash
# clone the repo
$ git clone https://github.com/rwajon/erc1155-events.git

$ cd erc1155-events

# run the development server
$ docker-compose up dev
```

After running the server, you can access API documentation at the following address `http://localhost:<PORT>/docs/index.html`

# Running tests

```bash
$ ./scripts/test.sh
```

### using docker & docker compose

```bash
$ docker-compose up test
```

# License

MIT
