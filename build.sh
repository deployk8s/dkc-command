#!/bin/bash

set -e
set -u

export GOPROXY=https://goproxy.cn

path=$(cd `dirname $0`;pwd)
rm -f dkc.tar.gz dkc.zip
mkdir -p dkc
rm -rf dkc/*

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dkc/dkc-linux-amd64 .
CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o dkc/dkc-darwin .
CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o dkc/dkc.exe .

cp -r README.md docs  inventory static views kubespray union_rsa dkc
#bee pack -be   GOOS=linux -v -exs=.go:.DS_Store:.tmp:cmd:conf:controllers:models:routers:src:tests:tmp:.tar:repo.json:.sh:dist -f=tar.gz -a=dkc-linux-amd64 -o dist
#bee pack -be   GOOS=darwin -v -exs=.go:.DS_Store:.tmp:cmd:conf:controllers:models:routers:src:tests:tmp:.tar:repo.json:.sh:dist -f=tar.gz -a=dkc-darwin -o dist
#bee pack -be   GOOS=windows -v -exs=.go:.DS_Store:.tmp:cmd:conf:controllers:models:routers:src:tests:tmp:.tar:repo.json:.sh:dist -f=zip -a=dkc -o dist
#bee pack -be   GOOS=windows -v -exs=.go:.DS_Store:.tmp:cmd:conf:controllers:models:routers:src:tests:tmp:.tar:repo.json:.sh:dist -f=tar.gz -a=dkc -o dist

#cd dist && zip dkc.zip dkc-darwin && zip dkc.zip dkc-linux-amd64

mkdir -p outputs
tar czf ./outputs/dkc.tar.gz dkc
zip -q -r ./outputs/dkc.zip dkc
