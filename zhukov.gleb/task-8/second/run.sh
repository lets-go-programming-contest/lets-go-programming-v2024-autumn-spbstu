#!/bin/bash

go build .\tags_main.go
./tags_main
go build -tags slice
./second

while true; do
    sleep 1
done