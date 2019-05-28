#!/bin/bash

export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

rm -rf $GOBIN/*
rm -rf $GOPATH/src/main/vendor
