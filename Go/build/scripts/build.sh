#!/usr/bin/env bash
cd $(dirname "$BASH_SOURCE")
cd ../../
GOOS=js GOARCH=wasm go build -o website/bin/main.wasm