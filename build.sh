#!/bin/bash

# shellcheck disable=SC2034
GOPROXY=https://goproxy.cn
cd console && npm install && npm run build
cd ../theme && npm install && npm run build
go build -i -v
echo 'build pipe done'
