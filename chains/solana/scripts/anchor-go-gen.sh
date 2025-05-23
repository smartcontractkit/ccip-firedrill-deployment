#!/usr/bin/env bash

set -e

for idl_path_str in "contracts/target/idl"/*
do
  IFS='/' read -r -a idl_path <<< "${idl_path_str}"
  IFS='.' read -r -a idl_name <<< "${idl_path[3]}"
  mkdir -p  ./gobindings/"${idl_name}"
  anchor-go -src "${idl_path_str}" -dst ./gobindings/"${idl_name}" -codec borsh
done

go fmt ./...
