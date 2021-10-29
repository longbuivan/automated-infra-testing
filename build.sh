#!/bin/bash

echo "starting localstack on docker"
docker-compose up -d


echo "running Unit Test for Terraform"
export GO111MODULE=on
cd test
go mod init terraform/test 
go mod tidy
go test -v -timeout 30m
