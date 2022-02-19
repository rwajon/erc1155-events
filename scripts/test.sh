#!/bin/sh

envs=""
rpc_url=$(grep -w .env -e 'TEST_RPC_URL' | sed 's/TEST_RPC_URL=//' | grep -v "#")

while IFS= read -r line; do
  env=$(echo $line | grep -v "#" | grep -v "GO_ENV" | grep -v "DEV" | grep -v "PROD")
  if [[ $env ]]; then
    envs="${envs}${env} "
  fi
done <"$(dirname "$0")/../.env"

echo $envs

for _ in {1..2}; do
  ./$(dirname "$0")/send_test_eth.sh --host=$rpc_url
done

rm $(dirname "$0")/../tmp/cover.out

go clean -testcache
eval "GO_ENV=test $envs go test ./tests... -v -coverpkg=./... -coverprofile=$(dirname "$0")/../tmp/cover.out $@"
go tool cover -html=$(dirname "$0")/../tmp/cover.out -o $(dirname "$0")/../tmp/cover.html
go tool cover -func=$(dirname "$0")/../tmp/cover.out
