#!/bin/bash

pushd .
cd nginx
docker build -t budget/nginx .
popd

pushd .
cd application
docker build --build-arg CACHE_DATE=$(date +%s) -t budget/core .
popd

