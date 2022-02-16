#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
  --amount=*)
    amount="${1#*=}"
    ;;
  --from=*)
    from="${1#*=}"
    ;;
  --to=*)
    to="${1#*=}"
    ;;
  *)
    printf "***************************\n"
    printf "* Error: Invalid argument (${1}).\n"
    printf "***************************\n"
    exit 1
    ;;
  esac
  shift
done

accounts=$(curl -X POST --data '{
            "jsonrpc":"2.0",
            "method":"personal_listAccounts",
            "params":[],
            "id":1
          }' -H "Content-Type: application/json" http://localhost:8545)

from_address=$(echo $accounts | python -c "import sys, json; print json.load(sys.stdin)['result'][0]")
to_address=$(echo $accounts | python -c "import sys, json; print json.load(sys.stdin)['result'][1]")

amount=${amount:-"1"}
from_address=${from:-$from_address}
to_address=${to:-$to_address}

if [[ $amount && $from_address && $to_address ]]; then
  echo "amount ==> $amount"
  echo "from_address ==> $from_address"
  echo "to_address ==> $to_address"

  curl -X POST --data '{
      "jsonrpc":"2.0",
      "method":"eth_sendTransaction",
      "params":[{"from":"'$from_address'", "to":"'$to_address'", "value":"'$amount'"}],
      "id":1
    }' -H "Content-Type: application/json" http://localhost:8545
fi
