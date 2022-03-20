#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
  --host=*)
    host="${1#*=}"
    ;;
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

host=${host:-"http://localhost:8545"}
amount=${amount:-"1"}

accounts=$(curl -X POST --data '{
            "jsonrpc":"2.0",
            "method":"personal_listAccounts",
            "params":[],
            "id":1
          }' -H "Content-Type: application/json" $host)

from_address=$(echo $(echo $accounts | awk -F'[:}]' '{print $(NF-1)}' | awk -F'"' '{print $(NF-1)}'))
to_address=$(echo $(echo $accounts | awk -F'[:}]' '{print $(NF-1)}' | awk -F'"' '{print $(NF-3)}'))
from_address=${from:-$from_address}
to_address=${to:-$to_address}

if [[ $amount && $from_address && $to_address ]]; then
  echo "amount: $amount | from: $from_address | to: $to_address"

  curl -X POST --data '{
      "jsonrpc":"2.0",
      "method":"eth_sendTransaction",
      "params":[{"from":"'$from_address'", "to":"'$to_address'", "value":"'$amount'"}],
      "id":1
    }' -H "Content-Type: application/json" $host
fi
