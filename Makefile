ORG                     := automatethethingsllc
TARGET_OS               := linux
TARGET_ARCH             := $(shell uname -m)

ARCH                    := $(shell go env GOARCH)
OS                      := $(shell go env GOOS)
LONG_BITS               := $(shell getconf LONG_BIT)

GOBIN                   := $(shell dirname `which go`)

PACKAGE                 := go-viper-cobra-skeleton
APPNAME                 := skeleton

APP_VERSION       		?= $(shell git describe --tags --abbrev=0)
GIT_TAG                 = $(shell git describe --tags)
GIT_HASH                = $(shell git rev-parse HEAD)
BUILD_DATE              = $(shell date '+%Y-%m-%d_%H:%M:%S')

ifeq ($(APP_VERSION),)
 	APP_VERSION = $(shell git branch --show-current)
endif

LDFLAGS=-X github.com/jeremyhahn/$(PACKAGE)/app.Name=${PACKAGE}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.Release=${APP_VERSION}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.GitHash=${GIT_HASH}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.GitTag=${GIT_TAG}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.BuildUser=${USER}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.BuildDate=${BUILD_DATE}
LDFLAGS+= -X github.com/jeremyhahn/$(PACKAGE)/app.Image=${IMAGE_NAME}

.PHONY: deps build build-debug build-static build-debug-static clean test initlog

default: build

deps:
	go get

build:
	$(GOBIN)/go build -o $(APPNAME) -ldflags="-w -s ${LDFLAGS}"

build-debug:
	$(GOBIN)/go build -gcflags='all=-N -l' -o $(APPNAME) -gcflags='all=-N -l' -ldflags="-w -s ${LDFLAGS}"

build-static:
	$(GOBIN)/go build -o $(APPNAME) --ldflags '-w -s -extldflags -static -v ${LDFLAGS}'

build-debug-static:
	$(GOBIN)/go build -o $(APPNAME) -gcflags='all=-N -l' --ldflags '-extldflags -static -v ${LDFLAGS}'

clean:
	$(GOBIN)/go clean
	rm -rf $(APPNAME) \
		$(APPNAME).log \
		/usr/local/bin/$(APPNAME)

test:
	cd app && $(GOBIN)/go test -v
	cd cmd && $(GOBIN)/go test -v

initlog:
	sudo touch /var/log/$(APPNAME).log && sudo chmod 777 /var/log/$(APPNAME).log