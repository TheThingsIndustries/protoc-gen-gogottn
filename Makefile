# Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

all: deps build

deps: 
	go mod vendor

build:
	go build .

.PHONY: deps build
