# {{cookiecutter.__copyright}}

# ==============================================================================
# Makefile for building Storage Central {{cookiecutter.service_name}} service
# ==============================================================================

# ------------------------------------------------------------------------------
# Configuration
#
# - These can be overwritten on invocation (e.g `make DIST_DIR=./out build`).
# ------------------------------------------------------------------------------

SHELL=/bin/bash -o pipefail

GO     := go
WIRE   := wire
DOCKER := docker

GCI       := gci
GOIMPORTS := goimports

export GOFLAGS     := -mod=vendor
export GOPRIVATE   := github.hpe.com

GOBUILD := $(GO) build
GOTEST  := $(GO) test

DIST_DIR     := ./dist
CMD_DIR      := ./cmd
INTERNAL_DIR := ./internal
MOCKS_DIR    := ./mocks
TEST_DIR     := ./test

# Mockery can cause a circular dependency with modules so disable them as they're not required
MOCKERY_CMD  := GOFLAGS="" mockery
MOCKERY_ARGS := --all --keeptree --case=snake --with-expecter --dir $(INTERNAL_DIR)

VERSION = $(shell (git describe --long --tags --match 'v[0-9]*' 2>/dev/null || echo v0.0.0) | cut -c2-)
LDFLAGS = -X main.Version=$(VERSION)
COMMIT  = $(shell git rev-parse --short HEAD)

UNIT_TEST_MODULES = $(shell $(GO) list $(INTERNAL_DIR)/...)
INTERNAL_GO_FILES = $(shell find $(INTERNAL_DIR) -type f -name '*.go')
INTERNAL_NON_TEST_GO_FILES = $(shell find $(INTERNAL_DIR) -type f -name '*.go' -not -name '*_test.go')
CMD_GO_FILES = $(shell find $(CMD_DIR) -type f -name '*.go')
ALL_GO_FILES = $(shell find . -type f -name '*.go' | grep -v '/vendor/' | grep -v '/mocks/')

# The CI_MINIMUM_TEST_COVERAGE environment variable is set automatically in
# Jenkins runs.  A default value is defined here for developer builds.
CI_MINIMUM_TEST_COVERAGE ?= 85

DOCKER_IMAGE := {{cookiecutter.service_name}}
DOCKER_TAG    = $(VERSION)

PUSH_REGISTRY := cds-harbor.rtplab.nimblestorage.com
PUSH_PROJECT  := sc-jenkins-branch
PUSH_IMAGE    := $(DOCKER_IMAGE_RDS_INVENTORY_MANAGER)
PUSH_TAG      := $(DOCKER_TAG)

# oci annotations to add to the image as labels
# these should technically be manifest annotations
# but that's currently experimental in the docker cli
# these can be viewed by inspecting the image
LABEL_CREATED   = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LABEL_AUTHORS  := {{cookiecutter.project_email}}
LABEL_SOURCE   := https://{{cookiecutter.__repository_url}}
LABEL_VERSION   = $(VERSION)
LABEL_REVISION  = $(COMMIT)
LABEL_VENDOR   := 'Hewlett Packard Enterprise'

# Build args for docker build
HTTPS_PROXY := ${HTTPS_PROXY}
HTTP_PROXY  := ${HTTP_PROXY}

# ------------------------------------------------------------------------------
# General Build Targets
# ------------------------------------------------------------------------------

.DELETE_ON_ERROR:

## help: Output this message and exit.
.PHONY: help
help:
	@grep -h '^##' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/## //'

## all: Run targets - checks, lint, and build.
.PHONY: all
all: checks lint build

## build: build the project
.PHONY: build
build: vendor
	# Since Mac uses nfs to mount the repositories, the user gets squashed down
	# to the underlying Mac user causing permission issues. The workaround is to
	# create a directory here which will be done within the docker container's
	# overlay filesystem as the Dockerfile calls the Makefile
	@mkdir -p dist
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(DIST_DIR) $(CMD_DIR)/...

## checks: Run any processes that might alter or generate code.
# Do fmt at the end to catch any generated code
.PHONY: checks
checks: tidy vendor wire mocks fmt

## tidy: Update GO modules (add missing and remove unused modules).
.PHONY: tidy
tidy:
	$(GO) mod tidy

## vendor: Download vendored dependencies.
vendor: go.mod go.sum
	$(GO) mod vendor

## wire: Generate dependency injection files using wire.
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/{{cookiecutter.service_name}}/...

## mocks: Generate the mocks for all internal interfaces for testing.
.PHONY: mocks
mocks: $(INTERNAL_NON_TEST_GO_FILES)
	rm -rf $(MOCKS_DIR)_maketemp/
	## Mockery returns error code 0 on these errors but produces incorrect output
	if $(MOCKERY_CMD) $(MOCKERY_ARGS) --output $(MOCKS_DIR)_maketemp 2>&1 | grep ERR; then \
		rm -rf $(MOCKS_DIR)_maketemp; \
		exit 1; \
	fi
	rm -rf $(MOCKS_DIR)/
	mv $(MOCKS_DIR)_maketemp $(MOCKS_DIR)

## fmt: Run the go formatters.
.PHONY: fmt
fmt: $(ALL_GO_FILES)
	gofmt -s -w -e $(ALL_GO_FILES)
	$(GCI) --local github.hpe.com -w $(ALL_GO_FILES)
	$(GOIMPORTS) -local github.hpe.com -w $(ALL_GO_FILES)

## dirty: Exits with non-zero exit code if there are any uncommitted changes.
## : Intended usage is within CI, checking that after rerunning code generation
## : commands the generated code is identical to the committed code.
.PHONY: dirty
dirty:
	if [ $$(git status --porcelain | wc -l) -ne "0" ]; then \
		echo "Missing / modified files:"; \
		git status --porcelain; \
		echo; \
		echo "Diff of changed files:"; \
		git diff; \
		exit 1; \
	fi

## clean: clean up built code and vendor directory
.PHONY: clean
clean:
	rm -rf $(DIST_DIR)

## lint: Run the project linters.
.PHONY: lint
lint: go-lint helmlint

## go-lint: Run golangci-lint for linter issues.
.PHONY: go-lint
go-lint: vendor
	@echo "Using config path:" $(shell golangci-lint config path) >&2
	# Display all issues, not just 50 per linter.
	golangci-lint run --timeout=10m --exclude vendor --max-issues-per-linter 0 --exclude '`atlantia` is a misspelling of `atlanta`'

## helmlint: Run the helm linter on all top level helm charts in the helm dir.
.PHONY: helmlint
helmlint:
	helm version
	helm lint $(shell find $(HELM_DIR) -mindepth 1 -maxdepth 1 -type d)

## trivy: Run trivy vulnerability scanner.
.PHONY: trivy
trivy:
	@echo 'NOTE: To suppress failures (with reason) see .trivyignore.'
	trivy fs --exit-code 1 --severity UNKNOWN,LOW,MEDIUM --no-progress .
	trivy fs --exit-code 1 --severity HIGH,CRITICAL --no-progress .

## test: Run all tests.
.PHONY: test
test: unit-test

## unit-test: Run all the unit tests.
.PHONY: unit-test
unit-test: vendor
	$(GOTEST) $(UNIT_TEST_MODULES) \
		-cover \
		-coverpkg=./... \
		-coverprofile={{cookiecutter.service_name}}.out \
		-count=1 \
		-v | tee {{cookiecutter.service_name}}-unit-test-output.txt
	cat {{cookiecutter.service_name}}.out | grep -v "/external/" | grep -v "/mocks/" > {{cookiecutter.service_name}}.cov

## code-coverage: Produce a per-function code coverage report from a unit test run.
.PHONY: code-coverage
code-coverage: {{cookiecutter.service_name}}.cov
	$(GO) tool cover -func {{cookiecutter.service_name}}.cov
	$(GO) tool cover -html={{cookiecutter.service_name}}.cov -o {{cookiecutter.service_name}}-unit-test-coverage.html
	$(eval TEST_COVERAGE = $(shell $(GO) tool cover -func {{cookiecutter.service_name}}.cov | grep 'total:' | awk '{print substr($$3, 1, length($$3)-1)}'))
	@echo "Unit tests passed with $(TEST_COVERAGE) coverage"
	$(eval PASSED_COVERAGE = $(shell awk 'BEGIN {printf ($(TEST_COVERAGE) < ${CI_MINIMUM_TEST_COVERAGE} ? "0" : "1")}'))
	@if [ $(PASSED_COVERAGE) == 0 ]; then echo "Require at least ${CI_MINIMUM_TEST_COVERAGE}% test coverage"; exit 1; fi

# ------------------------------------------------------------------------------
# Docker Targets
# ------------------------------------------------------------------------------

# tracking for built images
BUILT := .built

# helper functions
built          = $(DOCKER) images $(1) --format '{{cookiecutter._dot_ID}}' >> $(BUILT)
push           = $(DOCKER) push $(1)
pushed         = $(DOCKER) manifest inspect $(1) > /dev/null
push-immutable = $(call pushed,$(1)) || $(call push,$(1))

# Enable buildkit for better build features.
export DOCKER_BUILDKIT         := 1
# Don't collapse stage build progress once it completes.
export BUILDKIT_PROGRESS       := plain
# Enable experimental CLI for manifest support.
export DOCKER_CLI_EXPERIMENTAL := enabled

## docker-build: Build the docker images for this app.
.PHONY: docker-build
docker-build: docker-build-{{cookiecutter.service_name}}

## docker-build-{{cookiecutter.service_name}}: Build the {{cookiecutter.service_name}} image.
.PHONY: docker-build-{{cookiecutter.service_name}}
docker-build-{{cookiecutter.service_name}}:
	$(DOCKER) build \
		--pull \
		--force-rm \
		--target $(DOCKER_IMAGE) \
		--network host \
		--label org.opencontainers.image.created=$(LABEL_CREATED) \
		--label org.opencontainers.image.authors=$(LABEL_AUTHORS) \
		--label org.opencontainers.image.source=$(LABEL_SOURCE) \
		--label org.opencontainers.image.version=$(LABEL_VERSION) \
		--label org.opencontainers.image.revision=$(LABEL_REVISION) \
		--label org.opencontainers.image.vendor=$(LABEL_VENDOR) \
		--label org.opencontainers.image.title=$(DOCKER_IMAGE) \
		--build-arg HTTPS_PROXY=$(HTTPS_PROXY) \
		--build-arg HTTP_PROXY=$(HTTP_PROXY) \
		--build-arg VERSION=$(VERSION) \
		-t $(DOCKER_IMAGE):$(DOCKER_TAG) \
		.
	@$(call built,$(DOCKER_IMAGE):$(DOCKER_TAG))

## docker-push: Push the docker images for this app.
.PHONY: docker-push
docker-push: docker-push-{{cookiecutter.service_name}}

## docker-push-{{cookiecutter.service_name}}: Push the {{cookiecutter.service_name}} image.
.PHONY: docker-push-{{cookiecutter.service_name}}
docker-push-{{cookiecutter.service_name}}:
	$(DOCKER) tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(PUSH_REGISTRY)/$(PUSH_PROJECT)/$(PUSH_IMAGE):$(PUSH_TAG)
ifeq ($(IMMUTABLE_TAG),1)
	$(call push-immutable,$(PUSH_REGISTRY)/$(PUSH_PROJECT)/$(PUSH_IMAGE):$(PUSH_TAG))
else
	$(call push,$(PUSH_REGISTRY)/$(PUSH_PROJECT)/$(PUSH_IMAGE):$(PUSH_TAG))
endif

## docker-clean: Delete any built images.
.PHONY: docker-clean
docker-clean: $(BUILT)
	for image in $(shell cat $<); do $(DOCKER) rmi -f $$image || true; done
	rm -f $<

# ------------------------------------------------------------------------------
# Other Targets
# ------------------------------------------------------------------------------

build/certs/rds-ca-2019-root.pem:
	curl https://s3.amazonaws.com/rds-downloads/rds-ca-2019-root.pem  > ./build/certs/rds-ca-2019-root.pem

## config: Generate the deployment config for the service to use in developer environment
.PHONY: config
config:
	rm -f ./{{cookiecutter.service_name}}-deploy.yaml
	echo "# DO NOT EDIT, This file is autogenerated by 'make config'" > {{cookiecutter.service_name}}-deploy.yaml
	helm template -f helm/{{cookiecutter.service_name}}/dev-values.yaml {{cookiecutter.service_name}} helm/{{cookiecutter.service_name}}/ >> {{cookiecutter.service_name}}-deploy.yaml
