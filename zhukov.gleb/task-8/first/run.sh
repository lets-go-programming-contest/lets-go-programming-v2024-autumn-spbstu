#!/bin/bash

Remove-Item Env:\GOSSAFUNC
go tool compile -W main.go
go tool link main.o
$env:GOSSAFUNC = "main"
go tool compile main.go > ssa.html
go tool compile main.go
Remove-Item Env:\GOSSAFUNC

while true; do
    sleep 1
done