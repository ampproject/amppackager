#
# Copyright 2016-2020 The Libsacloud Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
TEST            ?=$$(go list ./... | grep -v vendor)
VETARGS         ?=-all
GOFMT_FILES     ?=$$(find . -name '*.go' | grep -v vendor)
AUTHOR          ?="The Libsacloud Authors"
COPYRIGHT_YEAR  ?="2016-2020"
COPYRIGHT_FILES ?=$$(find . -name "*.go" -print | grep -v "/vendor/")
export GO111MODULE=on

default: test vet

run:
	go run $(CURDIR)/main.go --disable-healthcheck $(ARGS)

test:
	go test ./sacloud $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-api: 
	go test ./api $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-builder:
	go test ./builder $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-utils: 
	go test ./utils/* $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-all: goimports vet test test-api test-builder test-utils

vet: golint
	go vet ./...

golint: 
	test -z "$$(golint ./... | grep -v 'vendor/' | grep -v '_string.go' | tee /dev/stderr )"

goimports: fmt
	goimports -l -w $(GOFMT_FILES)

fmt:
	gofmt -s -l -w $(GOFMT_FILES)


godoc:
	docker-compose up godoc

.PHONY: default test vet fmt golint test-api test-builder test-all run goimports

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get golang.org/x/lint/golint
	GO111MODULE=off go get github.com/motemen/gobump
	GO111MODULE=off go get github.com/sacloud/addlicense

.PHONY: bump-patch bump-minor bump-major version
bump-patch:
	gobump patch -w

bump-minor:
	gobump minor -w

bump-major:
	gobump major -w

version:
	gobump show

git-tag:
	git tag v`gobump show -r`

set-license:
	addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

