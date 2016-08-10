# Initialize version and gc flags
GO_LDFLAGS := -X `go list ./version`.GitCommit=`git rev-parse --short HEAD`
GO_GCFLAGS :=

# Full package list
PKGS := ./testconfig ./ov ./icsp ./liboneview ./rest ./utils

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

# clean: coverage-clean build-clean
# clean: coverage-clean dockerfile-clean
clean: dockerfile-clean
		echo "working on clean"

# check: dco fmt vet lint
check: dco fmt vet
validate: check
test: check test-short
# validate: check test-short test-long

# .PHONY: .all_build .all_coverage .all_release .all_test .all_validate test build validate clean
.PHONY: .all_coverage .all_release .all_test .all_validate test build validate clean
