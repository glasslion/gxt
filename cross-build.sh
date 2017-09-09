#!/bin/bash
rm -r ./builds/
env GOOS=windows GOARCH=386 go build -o ./builds/gtx-win-i386.exe
env GOOS=windows GOARCH=amd64 go build -o ./builds/gtx-win-adm64.exe
env GOOS=linux GOARCH=386 go build -o ./builds/gtx-linux-i386
env GOOS=linux GOARCH=amd64 go build -o ./builds/gtx-linux-amd64
env GOOS=linux GOARCH=arm go build -o ./builds/gtx-linux-arm
env GOOS=linux GOARCH=arm64 go build -o ./builds/gtx-linux-arm64
env GOOS=darwin GOARCH=386 go build -o ./builds/gtx-darwin-386
env GOOS=darwin GOARCH=arm64 go build -o ./builds/gtx-darwin-arm64
ls -la ./builds/



