# Cross builder helper
define glide-install
	glide install -s --update-vendored
endef

define glide-update
	glide update -s --update-vendored
endef


vendor-clean:
		@rm -rf $(PREFIX)/vendor/*
		@echo cleaning up in $(PREFIX)/vendor/*

# for fresh setup so we can get clean repo
glide-clean: vendor-clean
		@echo "Removing all glide data"
		rm -f $(PREFIX)/glide.lock
		@glide cache-clear

# setup the vendor folder with required packages that have been committed
glide-vendor:
		@$(call glide-install)
		@echo "Done placing packages into $(PREFIX)/vendor"

glide: glide-clean glide-vendor
		@echo "All done! run git status and commit to save any changes."

