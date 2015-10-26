# goal for this make file is to generate dockerfile from an upstream project

DOCKER_FILE ?=
DOCKER_FILE_URL ?=

PROXY_CONFIG_CONTENT ?= $(shell cat $(PROXY_DOCKER_ENV_FILE))

include mk/utils/proxy.mk

define dockerfile-clean
	[ -z "$(DOCKER_FILE)" ] || rm -f $(DOCKER_FILE);
endef

get-upstream-dockerfile:
		# get the dockerfile.machine from github.com/docker/machine
		curl -s $(DOCKER_FILE_URL) > $(DOCKER_FILE)

gen-dockerfile: proxy-config get-upstream-dockerfile
		echo 'setup proxy for $(DOCKER_FILE)'
		sed "/FROM.*/ r $(PROXY_DOCKER_ENV_FILE)" $(DOCKER_FILE) > $(DOCKER_FILE).t && mv $(DOCKER_FILE).t $(DOCKER_FILE)

dockerfile-clean:
		$(call dockerfile-clean)
