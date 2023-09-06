GOPATH:=$(shell go env GOPATH)

include scripts/make-rules/common.mk
include scripts/make-rules/format.mk
include scripts/make-rules/generate.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/help.mk

.PHONY: init
init: ## init devtools
	@pre-commit install


.DEFAULT_GOAL := help

