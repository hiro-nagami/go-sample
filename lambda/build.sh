#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o lambdatest lambda-test.go
