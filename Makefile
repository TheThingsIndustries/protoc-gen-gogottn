# Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

all: deps build

deps: 
	@command -v dep > /dev/null || go get github.com/golang/dep/cmd/dep
	dep ensure

build:
	go build .

.PHONY: deps build
