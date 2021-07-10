BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDDATE ?= $(shell date --iso-8601=seconds)
REVISION  ?= $(shell git rev-parse HEAD)
VERSION   ?= $(shell git log --date=short --pretty=format:'%h@%cd' -n 1 .)

GOOPTS ?=
ifneq (,$(wildcard vendor))
	GOOPTS := $(GOOPTS) -mod=vendor
endif

VERSION_LDFLAGS := \
  -X github.com/prometheus/common/version.Branch=$(BRANCH) \
  -X github.com/prometheus/common/version.BuildDate=$(BUILDDATE) \
  -X github.com/prometheus/common/version.Revision=$(REVISION) \
  -X github.com/prometheus/common/version.Version=$(VERSION)

all: test build

.PHONY: test
test:
	go test $(GOOPTS) ./...

build: amppkg

.PHONY: amppkg
amppkg:
	go build $(GOOPTS) -ldflags "$(VERSION_LDFLAGS)" -o amppkg ./cmd/amppkg/...
