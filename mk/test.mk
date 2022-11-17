# Quick test. You can bypass long tests using: `if testing.Short() { t.Skip("Skipping in short mode.") }`
TESTCONFIG_PACKAGE_ROOT_PATH := github.com/$(GH_USER)/$(GH_REPO)
TESTCONFIG_JSON_DATA_DIR := test/data
ONEVIEW_TEST_DATA=EGSL_HOUSTB200_LAB
# list of test case data can be included here
TEST_CASES ?= EGSL_HOUSTB200_LAB:$(HOME)/.oneview.houston.tb.200.env
TEST_RUN ?=

testcase = $(firstword $(subst :, ,$1))
testenv  = $(word 2,$(subst :, ,$1))

# acceptance test case helper
define goacceptance
	. $(call testenv,$(1)); \
	TESTCONFIG_PACKAGE_ROOT_PATH=$(TESTCONFIG_PACKAGE_ROOT_PATH) \
	TESTCONFIG_JSON_DATA_DIR=$(TESTCONFIG_JSON_DATA_DIR) \
	ONEVIEW_TEST_DATA=$(call testcase,$(1)) \
	ONEVIEW_TEST_ACCEPTANCE=true \
	ICSP_TEST_ACCEPTANCE=true \
	$(GO) test $(VERBOSE_GO) \
	  -test.timeout=60m -test.v=true --short \
		-tags "$(BUILDTAGS)" $(PKGS) $(TEST_RUN);
endef

test-short:
	TESTCONFIG_PACKAGE_ROOT_PATH=$(TESTCONFIG_PACKAGE_ROOT_PATH) \
	TESTCONFIG_JSON_DATA_DIR=$(TESTCONFIG_JSON_DATA_DIR) \
	ONEVIEW_TEST_DATA=$(ONEVIEW_TEST_DATA) \
	$(GO) test $(VERBOSE_GO) -test.short -tags "$(BUILDTAGS)" $(PKGS) $(TEST_RUN)

# Runs long tests also, plus race detection
test-long:
	TESTCONFIG_PACKAGE_ROOT_PATH=$(TESTCONFIG_PACKAGE_ROOT_PATH) \
	TESTCONFIG_JSON_DATA_DIR=$(TESTCONFIG_JSON_DATA_DIR) \
	ONEVIEW_TEST_DATA=$(ONEVIEW_TEST_DATA) \
	$(GO) test $(VERBOSE_GO) -race -tags "$(BUILDTAGS)" $(PKGS) $(TEST_RUN)

# Runs acceptance test, requires a connection to real system
test-acceptance:
	$(foreach TEST_CASE_DATA,$(TEST_CASES),$(call goacceptance,$(TEST_CASE_DATA),$(TEST_ENV),$<))

#TODO: non-functional, determine if we really need this for the library
test-integration:
	TESTCONFIG_PACKAGE_ROOT_PATH=$(TESTCONFIG_PACKAGE_ROOT_PATH)
	TESTCONFIG_JSON_DATA_DIR=$(TESTCONFIG_JSON_DATA_DIR)
	ONEVIEW_TEST_DATA=$(ONEVIEW_TEST_DATA)
	$(eval TESTSUITE=$(filter-out $@,$(MAKECMDGOALS)))
	test/integration/run-bats.sh $(TESTSUITE)

%:
	@:
