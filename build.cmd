@echo off
@echo Building windows executable
@go build cmf.go
echo Building OSX executable

@set GOOS=darwin
@go build cmf.go
@set GOOS=windows
@echo on