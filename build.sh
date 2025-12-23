#!/bin/sh

go \
	build \
	-v \
	-o ./cmd/timeout/timeout \
	./cmd/timeout
