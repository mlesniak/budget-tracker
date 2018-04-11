#!/bin/bash

pushd .
cd nginx
docker build -t budget/nginx .
popd

pushd .
cd application
docker build -t budget/core .
popd

