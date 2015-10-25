# goal for this make file is to generate dockerfile from an upstream project

DOCKER_FILE ?=
DOCKER_FILE_URL ?=

PROXY_CONFIG_CONTENT ?= $(shell cat $(PROXY_DOCKER_ENV_FILE))

include mk/utils/proxy.mk

gen-dockerfile: proxy-config
		echo 'setup proxy for $(DOCKER_FILE)'
		sed "/FROM.*/ r $(PROXY_DOCKER_ENV_FILE)" $(DOCKER_FILE) > $(DOCKER_FILE).t && mv $(DOCKER_FILE).t $(DOCKER_FILE)
