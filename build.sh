#!/bin/bash
export GOPATH=$HOME/go
if [ "${GITHUB_ACTIONS}" == "true" ]; then
go test 1> debug.out
else
go test -v
fi

go get 1>> debug.out

BINARY="lunar"

if [ $? == 0 ]; then
  if [ "${GOOS}" == "windows" ]; then
    if [ "${GITHUB_ACTIONS}" == "true" ]; then
go build -v -ldflags="-X main.gitver=$(git describe --always --long --dirty)" -o ${BINARY}.exe *.go 1>> debug.out
echo "${BINARY}.exe"
    else
go build -v -ldflags="-X main.gitver=$(git describe --always --long --dirty)" -o ${BINARY}.exe *.go
    fi
  else
    if [ "${GITHUB_ACTIONS}" == "true" ]; then
go build -v -ldflags="-X main.gitver=$(git describe --always --long --dirty)" -o ${BINARY} *.go 1>> debug.out
echo "${BINARY}"
    else
go build -v -ldflags="-X main.gitver=$(git describe --always --long --dirty)" -o ${BINARY} *.go
    fi
  fi
fi