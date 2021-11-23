#!/usr/bin/env bash

go build -ldflags="-s -w" .
upx example > /dev/null 2>&1
