# Makefile for Go project status check

.PHONY: check vet

check: vet

vet:
	echo "Running go vet"
	go vet ./...