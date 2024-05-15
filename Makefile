ORG                     := automatethethingsllc
TARGET_OS               := linux
TARGET_ARCH             := $(shell uname -m)

ARCH                    := $(shell go env GOARCH)
OS                      := $(shell go env GOOS)
LONG_BITS               := $(shell getconf LONG_BIT)

GOBIN                   := $(shell dirname `which go`)

APP                     := go-viper-cobra-skeleton

APP_VERSION       		?= $(shell git describe --tags --abbrev=0)
GIT_TAG                 = $(shell git describe --tags)
GIT_HASH                = $(shell git rev-parse HEAD)
BUILD_DATE              = $(shell date '+%Y-%m-%d_%H:%M:%S')

ifeq ($(APP_VERSION),)
 	APP_VERSION = $(shell git branch --show-current)
endif

LDFLAGS=-X github.com/jeremyhahn/$(APP)/app.Name=${APP}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.Release=${APP_VERSION}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.GitHash=${GIT_HASH}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.GitTag=${GIT_TAG}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.BuildUser=${USER}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.BuildDate=${BUILD_DATE}
LDFLAGS+= -X github.com/jeremyhahn/$(APP)/app.Image=${IMAGE_NAME}

.PHONY: deps build build-debug build-static build-debug-static clean test initlog

default: build

deps:
	go get

build:
	$(GOBIN)/go build -o $(APP) -ldflags="-w -s ${LDFLAGS}"

build-debug:
	$(GOBIN)/go build -gcflags='all=-N -l' -o $(APP) -gcflags='all=-N -l' -ldflags="-w -s ${LDFLAGS}"

build-static:
	$(GOBIN)/go build -o $(APP) --ldflags '-w -s -extldflags -static -v ${LDFLAGS}'

build-debug-static:
	$(GOBIN)/go build -o $(APP) -gcflags='all=-N -l' --ldflags '-extldflags -static -v ${LDFLAGS}'

clean:
	$(GOBIN)/go clean
	rm -rf $(APP) \
		$(APP).log \
		/usr/local/bin/$(APP)

test:
	cd app && $(GOBIN)/go test -v
	cd cmd && $(GOBIN)/go test -v

initlog:
	sudo touch /var/log/$(APP).log && sudo chmod 777 /var/log/$(APP).log