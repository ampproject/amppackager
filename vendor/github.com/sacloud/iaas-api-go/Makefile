#====================
AUTHOR         ?= The sacloud/iaas-api-go Authors
COPYRIGHT_YEAR ?= 2022-2023

include includes/go/common.mk
#====================

default: $(DEFAULT_GOALS)

.PHONY: tools
tools: dev-tools

.PHONY: clean-all
clean-all:
	find . -type f -name "zz_*.go" -delete

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen:
	go generate ./...