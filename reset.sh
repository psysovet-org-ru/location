#!/usr/bin/env sh
export $(grep -v '^#' .env | xargs)

./build/location migrate down 1
./build/location migrate apply