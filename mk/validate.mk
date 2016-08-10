# Validate DCO on all history
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

# XXX vendorized script miss exec bit, hence the gymnastic
# plus the path resolution...
# TODO migrate away from the shell script and have a make equivalent instead
dco:
	@echo 'Performing DCO checks'
	@echo `bash $(current_dir)/../build/validate-dco`

# Fmt
fmt:
	@echo 'Performing FMT checks, if any files appear they run gofmt -s -w on each file'
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v vendor/ | grep -v Vendor/ | tee /dev/stderr)"

# Vet
vet: build
	@echo 'Performing VET checks'
	@echo $(PKGS)
	@test -z "$$(go vet $(PKGS) 2>&1 | tee /dev/stderr)"

# Lint
lint:
	@echo 'Performing lint checks'
	$(if $(GOLINT), , \
		$(error Please install golint: go get -u github.com/golang/lint/golint))
	@test -z "$$($(GOLINT) ./... 2>&1 | grep -v vendor/ | grep -v Vendor/ | tee /dev/stderr)"
