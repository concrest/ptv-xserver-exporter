@echo off
SET GO111MODULE=on
go test
go build -ldflags "-X 'main.version=%BUILD_BUILDNUMBER%' -X 'main.buildDate=%date%T%time%' -X 'main.commitHash=%BUILD_SOURCEVERSION%'"