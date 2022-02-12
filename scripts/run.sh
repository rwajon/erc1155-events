#!/bin/bash

port=""

if [[ -f .env ]]; then
  port=$(grep -w .env -e 'PORT' | sed 's/PORT=//' | grep -v "#")
fi
while true; do
  if [[ $port ]]; then
    pid=$(lsof -t -i :$port)
  fi
  if [[ $pid ]]; then
    echo -e "Port $port address already in use.\nPID: $pid\nEnter a new port: "
    read answer
    port=$answer
    pid=""
  else
    break
  fi
done

if ! [[ -f $(go env GOPATH)/bin/air ]]; then
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

PORT=$port $(go env GOPATH)/bin/air
