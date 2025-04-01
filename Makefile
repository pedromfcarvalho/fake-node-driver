all: build

.PHONY: build
build:
	go build -ldflags "-w -s" -o docker-machine-driver-fake
