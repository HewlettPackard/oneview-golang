# # Plain make targets if not requested inside a container

include Makefile.inc

ifneq (,$(findstring test-integration,$(MAKECMDGOALS)))
	include mk/main.mk
else ifeq ($(USE_CONTAINER),)
	include mk/main.mk
else
# Otherwise, with docker, swallow all targets and forward into a container
DOCKER_IMAGE_NAME := "oneview-golang-build"
DOCKER_CONTAINER_NAME := oneview-golang-build-container
# get the dockerfile from docker/machine project so we stay in sync with the versions they use for go
DOCKER_FILE := Dockerfile


build: gen-dockerfile
test: build
%:
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
		    $(DOCKER_IMAGE_NAME) \
		    make $@

endif

include mk/utils/dockerfile.mk
include mk/utils/godeps.mk
