#!/bin/sh

go clean -testcache
GO_ENV=test go test ./tests... ${@}
