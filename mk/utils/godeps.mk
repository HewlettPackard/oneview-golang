# setup any dependencies for Build

# GOPATH := $(HOME)/go
# PATH := $(PATH):$(GOPATH)/bin:/usr/local/go/bin
GO_PACKAGES := github.com/docker/machine github.com/stretchr/testify/assert
GO15VENDOREXPERIMENT := 1

# Cross builder helper
define godeps-get
	GOPCKG=$(1);
	godep get $(GOPCKG);
endef

define godeps-save
	godep save $(1);
endef

define godeps-clean
	echo 'Clean Package $(1)';
	[ -d $(GOPATH)/src/$(1) ] && \
		( cd $(GOPATH)/src/$(1); \
			_PKG_CLEAN=`git rev-parse --show-toplevel`; \
			[ -d $$_PKG_CLEAN ] && rm -rf $$_PKG_CLEAN; ) || \
		echo "Skipting clean for $(1)";
endef

vendor-clean:
		@rm -rf $(PREFIX)/vendor
		@echo cleaning up in $(PREFIX)/vendor

# for fresh setup so we can do godep save -r
godeps-clean:
		@echo "Removing all dependent packages from $(GOPATH)"
		$(foreach GOPCKG,$(GO_PACKAGES),$(call godeps-clean,$(GOPCKG)))
		rm -rf $(GOPATH)/src/github.com/$(GH_USER)/$(GH_REPO)

# setup a fresh GOPATH directory with what would be needed to build
godeps-init: godeps-clean
		@echo "Pulling required packages into $(GOPATH)"
		mkdir -p $(GOPATH)/src/github.com/$(GH_USER)
		ln -s $(PREFIX) $(GOPATH)/src/github.com/$(GH_USER)/$(GH_REPO)
		@echo "Get dependent packages"
		$(foreach GOPCKG,$(GO_PACKAGES),$(call godeps-get,$(GOPCKG)))
		$(call godeps-save,$(GO_PACKAGES))

godeps: godeps-init
godep: godeps
