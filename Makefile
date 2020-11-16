.PHONY: all

VERSION := 1.0

all: build

build:
	docker build --rm -t liquid-cdt .