# use PROXY_CONFIG

PROXY_DOCKER_ENV_FILE := .proxy.docker.env

define DOCKER_HTTPS_PROXY
\nENV HTTPS_PROXY %s\n
endef
define DOCKER_HTTP_PROXY
ENV HTTP_PROXY %s\n
endef
define DOCKER_https_proxy
ENV https_proxy %s\n
endef
define DOCKER_http_proxy
ENV http_proxy %s\n
endef
define DOCKER_NO_PROXY
ENV NO_PROXY %s\n
endef
define DOCKER_no_proxy
ENV no_proxy %s\n
endef

ifdef HTTPS_PROXY
DOCKER_PROXY_CONFIG := $(DOCKER_HTTPS_PROXY)
PROXY_CONFIG := '$(HTTPS_PROXY)'
endif
ifdef HTTP_PROXY
DOCKER_PROXY_CONFIG := $(DOCKER_PROXY_CONFIG)$(DOCKER_HTTP_PROXY)
PROXY_CONFIG := $(PROXY_CONFIG) '$(HTTP_PROXY)'
endif
ifdef https_proxy
DOCKER_PROXY_CONFIG := $(DOCKER_PROXY_CONFIG)$(DOCKER_https_proxy)
PROXY_CONFIG := $(PROXY_CONFIG) '$(https_proxy)'
endif
ifdef http_proxy
DOCKER_PROXY_CONFIG := $(DOCKER_PROXY_CONFIG)$(DOCKER_http_proxy)
PROXY_CONFIG := $(PROXY_CONFIG) '$(http_proxy)'
endif
ifdef NO_PROXY
DOCKER_PROXY_CONFIG := $(DOCKER_PROXY_CONFIG)$(DOCKER_NO_PROXY)
PROXY_CONFIG := $(PROXY_CONFIG) '$(NO_PROXY)'
endif
ifdef no_proxy
DOCKER_PROXY_CONFIG := $(DOCKER_PROXY_CONFIG)$(DOCKER_no_proxy)
PROXY_CONFIG := $(PROXY_CONFIG) '$(no_proxy)'
endif

proxy-clean:
		rm -f $(PROXY_DOCKER_ENV_FILE)

ifneq ($(PROXY_CONFIG),)
proxy-config: proxy-clean
		# generate a
		printf '$(DOCKER_PROXY_CONFIG)' $(PROXY_CONFIG) > $(PROXY_DOCKER_ENV_FILE)
else
proxy-config: proxy-clean
		touch $(PROXY_DOCKER_ENV_FILE)
endif
