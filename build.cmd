@echo off

@echo Building windows executable
@set GOOS=windows
@set GOARCH=amd64
@go build cmf.go

echo Building OSX executable
@set GOOS=darwin
@go build cmf.go

REM set back GOOS to be windows
@set GOOS=windows

@echo on