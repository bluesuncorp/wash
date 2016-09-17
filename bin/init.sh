#!/bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"

echo "Script Running From $DIR"

ROOT=$DIR/../..

echo "Moving to ROOT (${DIR}) directory"
cd $ROOT

echo "Running go get"
# imports needed packages
go get -u
go get

echo "Installing Compile Daemon"
go get -u github.com/go-playground/justdoit