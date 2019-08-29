#!/bin/bash

export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

cd $GOPATH/src/main
glide install
cd $GOPATH

go install $GOPATH/src/main
