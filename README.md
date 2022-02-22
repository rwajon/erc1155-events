# ERC1155-events

ERC1155-events is an application that subscribes to smart contract transfer events and store them.
To get transfer events from an Ethereum smart contract address, you have to first add it to a watch list.
To add an address in the watch list, please refer to the the [API documentation](https://erc1155-events.herokuapp.com/docs/index.html)

- [Live API URL](https://erc1155-events.herokuapp.com)
- [API documentation](https://erc1155-events.herokuapp.com/docs/index.html)

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
