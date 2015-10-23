#!/bin/bash
USAGEERR=2
[[ $_ != $0 ]] && USAGEERR=1 || USAGEERR=0
if [ $USAGEERR -eq 1 ]; then
  echo "USAGE: source $0"
  echo ""
  echo "This script must be sourced"
  exit 1
fi

TEMP_DIR=$(dirname $(mktemp -u))
SCRIPT_HOME="$(dirname $0)"
CURRENT_DIR=$(pwd)
# cd to the root folder of this repo
# ie;./drivers/oneview/scripts/conf
cd "$(git rev-parse --show-toplevel)"
echo "Working repo : $(pwd)"
printf "Please select a profile to source:\n"
select oneviewenv in $HOME/.oneview*.env; do test -n "$oneviewenv" && break; echo ">>> Invalid Selection"; done
. $oneviewenv

#
# setup container env file
#
echo "setting up container env file"
cat $oneviewenv |grep -v '.*#.*'|sed 's/export //g' > $SCRIPT_HOME/.oneview.env


#
# install docker-compose
#
echo "Checking for docker-compose, otherwise download it"
$TEMP_DIR/docker-compose --version 2>&1 > /dev/null || \
docker-compose --version 2>&1 > /dev/null || \
    ( echo "Get current version"; \
    VERSION_NUM=$(curl -s https://github.com/docker/compose/releases/latest|awk -F'href="' '{print $2}'|sed 's/">.*//g'|sed 's/.*tag\///g');\
    OS=$(uname -s);\
    PLATFORM=$(uname -m);\
    curl -L https://github.com/docker/compose/releases/download/$VERSION_NUM/docker-compose-$OS-$PLATFORM > $TEMP_DIR/docker-compose && \
    chmod +x $TEMP_DIR/docker-compose)

export DOCKER_COMPOSE_BIN=$TEMP_DIR/docker-compose
$DOCKER_COMPOSE_BIN --version
echo $?


#
# setup compose file
#
COMPOSE_FILE=$SCRIPT_HOME/conf/$(echo $ONEVIEW_TEST_DATA | tr '[:upper:]' '[:lower:]').yaml
if [ ! -f $COMPOSE_FILE ]; then
    echo "Could not find $COMPOSE_FILE to start"
    return 1
fi


#
# proxy
#
rm -f $SCRIPT_HOME/.proxy.env
touch $SCRIPT_HOME/.proxy.env
if [ ! "$https_proxy" ] ; then
echo "setup proxy"
cat > $SCRIPT_HOME/.proxy.env << PROXY
HTTP_PROXY="$https_proxy"
HTTPS_PROXY="$https_proxy"
NO_PROXY="$no_proxy"
PROXY
fi

#
# compose
#
$DOCKER_COMPOSE_BIN -f $COMPOSE_FILE rm -f -v
$DOCKER_COMPOSE_BIN -f $COMPOSE_FILE up -d

#
# setup aliases
#
X_ALIAS=$(cat $COMPOSE_FILE|grep container_name|awk -F':' '{print $2}'|sed 's/\s//g' | \
    xargs -i echo alias {}=\'docker exec {} godep go test -test.timeout=60m -test.v=true --short\')
eval $X_ALIAS
echo ""
echo "Sample Usage:"
echo ""
cat $COMPOSE_FILE|grep container_name|awk -F':' '{print $2}'|sed 's/\s//g' | \
    xargs -i echo {} ./drivers/oneview/ov

cd "${CURRENT_DIR}"
return 0
