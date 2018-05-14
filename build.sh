#!/bin/bash
#server Name
binName=web-demo
gopath=$(dirname $(dirname `pwd`))
export GOPATH=$gopath

#编译
go build -o release/bin/$binName