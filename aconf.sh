#!/bin/bash

ACONF_BIN="/usr/local/bin/aconf"
output=$($ACONF_BIN "$@")

if [ $? -ne 0 ]; then
  echo "$output"
  exit 1
fi

eval "$output"
