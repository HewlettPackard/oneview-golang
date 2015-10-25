# setup any dependencies for Build

# GOPATH := $(HOME)/go
# PATH := $(PATH):$(GOPATH)/bin:/usr/local/go/bin
GO15VENDOREXPERIMENT := 1

vendor-clean:
		@rm -rf $(PREFIX)/vendor
		@echo cleaning up in $(PREFIX)/vendor

# for fresh setup so we can do godep save -r
godeps-clean:
		@echo "Removing all dependent packages from $(GOPATH)"
		rm -rf $(GOPATH)/src/github.com/docker/machine
		rm -rf $(GOPATH)/src/github.com/HewlettPackard/oneview-golang
		rm -rf $(GOPATH)/src/github.com/$(GH_USER)/$(GH_REPO)

# setup a fresh GOPATH directory with what would be needed to build
godeps-init: godeps-clean
		@echo "Pulling required packages into $(GOPATH)"
		mkdir -p $(GOPATH)/src/github.com/$(GH_USER)
		ln -s $(PREFIX) $(GOPATH)/src/github.com/$(GH_USER)/$(GH_REPO)
		@echo "Get dependent packages"
		godep get github.com/docker/machine
		godep save github.com/docker/machine
		#TODO : change this once we release and remove the key
		if [ -d $(PREFIX)/../oneview-golang/.git ];
				ln -s $(PREFIX)/../oneview-golang $(GOPATH)/src/github.com/$(GH_USER)/oneview-golang
		else
				git clone git@github.com:HewlettPackard/oneview-golang.git $(PREFIX)/../oneview-golang
				ln -s $(PREFIX)/../oneview-golang $(GOPATH)/src/github.com/$(GH_USER)/oneview-golang
		fi
		# godep get github.com/HewlettPackard/oneviwe-golang/
		# godep save github.com/HewlettPackard/oneviwe-golang/
