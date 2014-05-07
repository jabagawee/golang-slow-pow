#! /bin/bash

go test -bench . -cpuprofile test_profile
go test -c
mv *.test test_binary
go tool pprof test_binary test_profile
