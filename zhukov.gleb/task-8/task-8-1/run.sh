#!/bin/bash

go tool compile -W cmd/main.go > AST.txt
go tool link main.o

while true; do
    sleep 1
done