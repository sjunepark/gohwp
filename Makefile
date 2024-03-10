# Makefile for Go project status check

.PHONY: check vet

check: vet fmt

vet:
	echo "Running go vet"
	go vet ./...

fmt:
	echo "Running go fmt"
	go fmt ./...
