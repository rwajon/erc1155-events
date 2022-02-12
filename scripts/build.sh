#!/bin/bash

while [ $# -gt 0 ]; do
  case "$1" in
  --os=*)
    os="${1#*=}"
    ;;
  --arch=*)
    arch="${1#*=}"
    ;;
  --app_name=*)
    app_name="${1#*=}"
    ;;
  *)
    printf "***************************\n"
    printf "* Error(get_args.sh): Invalid argument (${1}).\n"
    printf "***************************\n"
    exit 1
    ;;
  esac
  shift
done

os=${os:-"linux"}
arch=${arch:-"amd64"}
app_name=${app_name:-"app"}

echo "env GOOS=$os GOARCH=$arch go build -o ${app_name}-${arch} ."
env GOOS=$os GOARCH=$arch go build -o ${app_name}-${arch} .
mkdir -p bin
mv ${app_name}-${arch} bin/
