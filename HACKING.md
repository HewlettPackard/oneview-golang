Acceptance test cases
---------------------
This guide will describe steps we're using on running acceptance test cases.

Environment setup
------------------
Typical scenario, `make build` and `make test` should work as normal for the
overall project.

craete an environment, `$HOME/.oneview.houston.tb.200.env`, script to export these values:

```bash
cat > "$HOME/.oneview.env" << ONEVIEW
export ONEVIEW_APIVERSION=120

export ONEVIEW_ILO_USER=docker
export ONEVIEW_ILO_PASSWORD=password

export ONEVIEW_ICSP_ENDPOINT=https://15.x.x.x
export ONEVIEW_ICSP_USER=username
export ONEVIEW_ICSP_PASSWORD=password
export ONEVIEW_ICSP_DOMAIN=LOCAL

export ONEVIEW_OV_ENDPOINT=https://15.x.x.x
export ONEVIEW_OV_USER=username
export ONEVIEW_OV_PASSWORD=password
export ONEVIEW_OV_DOMAIN=LOCAL

export ONEVIEW_SSLVERIFY=true

ONEVIEW

```
Now you can setup environment value for the test cases you plan to run.

```bash
export TEST_CASES=EGSL_HOUSTB200_LAB:~/.oneview.houston.tb.200.env
```
Run the acceptance test
-------------------------
Acceptance test can be executed:

```bash
make test-acceptance
```

Running debug log output
-------------------------
Output from test case debugging log can be handy.

```bash
ONEVIEW_DEBUG=true make test-acceptance
```

Run a single specific test
---------------------------
Sometimes it's usefull to run just a single test case.
```bash
ONEVIEW_DEBUG=true make test-acceptance TEST_RUN='-test.run=TestGetAPIVersion'
```
