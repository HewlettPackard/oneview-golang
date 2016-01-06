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

Updating external dependencies
------------------------------
This project relies on external libraries that are committed into Godeps folder.
The libraries are used throughout the project to maintain compatibility with
projects such as docker-machine-oneview and docker/machine project.

1. Start by cleaning all libraries from Godeps folder

   ```
   make godeps-clean
   ```
   You can verify the execution by checking that there are no files left in the folder `Godeps/_workspace/src`.

2. Get the latest packages with godeps target

   ```
   make godeps
   ```

3. Run a build in a docker container.

   ```
   USE_CONTAINER=true make test
   ```

4. Evaluate changes.
   At this point you might have changes to the dependent libraries that have to be incorporated into the build process.   Update any additional or no longer libraries by editing the file : [mk/utils/godeps.mk](mk/utils/godeps.mk).  This file contains arguments GO_PACKAGES that should have a space separated list of all needed packages.
   Whenever adjusting libraries, make sure to re-do steps 1-3 iteratively.

5. Ok, it all test and passes, so it's time to commit your changes.

  ```
  git add --all
  ```
  Use `git status` to review additions, removals, and changes.
  Use `git commit -s -m "library update version X.X"` to commit your changes.
