# # Plain make targets if not requested inside a container

USE_CONTAINER ?= true

define noop_targets
	@make -pn | sed -rn '/^[^# \t\.%].*:[^=]?/p'|grep -v '='| grep -v '(%)'| grep -v '/'| awk -F':' '{print $$1}'|sort -u;
endef

include Makefile.inc

ifneq (,$(findstring test-integration,$(MAKECMDGOALS)))
	include mk/main.mk
else ifeq ($(USE_CONTAINER),false)
	include mk/main.mk
else
# Otherwise, with docker, swallow all targets and forward into a container
DOCKER_IMAGE_NAME := "oneview-golang-build"
DOCKER_CONTAINER_NAME := oneview-golang-build-container
# get the dockerfile from docker/machine project so we stay in sync with the versions they use for go
DOCKER_FILE_URL := file://$(PREFIX)/Dockerfile
DOCKER_FILE := .dockerfile.oneview


noop:
	@echo When using 'USE_CONTAINER' use a "make <target>"
	@echo
	@echo Possible targets
	@echo
	$(call noop_targets)

clean: gen-dockerfile
test: gen-dockerfile
%:
		export GO15VENDOREXPERIMENT=1
		docker build -f $(DOCKER_FILE) -t $(DOCKER_IMAGE_NAME) .

		test -z '$(shell docker ps -a | grep $(DOCKER_CONTAINER_NAME))' || docker rm -f $(DOCKER_CONTAINER_NAME)

		docker run --name $(DOCKER_CONTAINER_NAME) \
		    -e DEBUG \
		    -e STATIC \
		    -e VERBOSE \
		    -e BUILDTAGS \
		    -e PARALLEL \
		    -e COVERAGE_DIR \
		    -e TARGET_OS \
		    -e TARGET_ARCH \
		    -e PREFIX \
		    -e GO15VENDOREXPERIMENT \
		    -e TEST_RUN \
		    -e ONEVIEW_DEBUG \
		    -e GH_USER \
		    -e GH_REPO \
		    -e USE_CONTAINER=false \
		    $(DOCKER_IMAGE_NAME) \
		    make $@

endif

include mk/utils/dockerfile.mk
include mk/utils/glide.mk
