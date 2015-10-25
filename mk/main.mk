# Initialize version and gc flags
GO_LDFLAGS := -X `go list ./version`.GitCommit=`git rev-parse --short HEAD`
GO_GCFLAGS :=

# Full package list
# PKGS := $(shell go list -tags "$(BUILDTAGS)" ./testconfig/... ./ov/... ./icsp/... ./liboneview/... ./rest/... ./utils/... | grep -v "/vendor/" | grep -v "/Godeps/")
PKGS := ./testconfig ./ov ./icsp ./liboneview ./rest ./utils

# Support go1.5 vendoring (let us avoid messing with GOPATH or using godep)
export GO15VENDOREXPERIMENT = 1

# Resolving binary dependencies for specific targets
GOLINT_BIN := $(GOPATH)/bin/golint
GOLINT := $(shell [ -x $(GOLINT_BIN) ] && echo $(GOLINT_BIN) || echo '')

# Honor debug
ifeq ($(DEBUG),true)
	# Disable function inlining and variable registerization
	GO_GCFLAGS := -gcflags "-N -l"
else
	# Turn of DWARF debugging information and strip the binary otherwise
	GO_LDFLAGS := $(GO_LDFLAGS) -w -s
endif

# Honor static
ifeq ($(STATIC),true)
	# Append to the version
	GO_LDFLAGS := $(GO_LDFLAGS) -extldflags -static
endif

# Honor verbose
VERBOSE_GO :=
GO := go
ifeq ($(VERBOSE),true)
	VERBOSE_GO := -v
	GO := go
endif

# include mk/build.mk
# include mk/coverage.mk
# include mk/release.mk
include mk/test.mk
include mk/validate.mk

# TODO: what to do with build/release/coverage
# .all_build: build build-clean build-x build-machine build-plugins
# .all_coverage: coverage-generate coverage-html coverage-send coverage-serve coverage-clean
# .all_release: release-checksum release
# TODO: work on test-long and test-integration
# .all_test: test-short test-long test-integration
.all_test: test-short
.all_validate: dco fmt vet lint

default: test
# Build native machine and all drivers
# TODO: cleanup build: build-machine build-plugins
# build: build-x

#TODO: cleanup
# Just build native machine itself
# machine: build-machine
# Just build the native plugins
# plugins: build-plugins
# Build all, cross platform
# cross: build-x

# clean: coverage-clean build-clean
clean: coverage-clean
# check: dco fmt vet lint
check: dco fmt
validate: check
test: check test-short
# validate: check test-short test-long

# .PHONY: .all_build .all_coverage .all_release .all_test .all_validate test build validate clean
.PHONY: .all_coverage .all_release .all_test .all_validate test build validate clean
