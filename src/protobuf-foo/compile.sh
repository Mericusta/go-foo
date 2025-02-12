#!/bin/bash

rm -rf ./pb/*.go

protoc --proto_path=proto --go_out=. proto/dungeon.proto
