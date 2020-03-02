#!/bin/bash

# import
DIR=$(cd $(dirname $0) && pwd)
. ${DIR}/base.sh

# make build directory
mkdir -p ${BUILD_DIR}
rm -Rf ${BUILD_DIR}/*

BUILD_DATE=$(date '+%Y/%m/%d %H:%M:%S %Z')

# build api
CGO_LDFLAGS="`mecab-config --libs`" \
CGO_FLAGS="`mecab-config --inc-dir`" \
go build \
-o ${BUILD_DIR}/${API_BIN} \
-ldflags="-X main.branch=${BRANCH} -X main.revision=${LONG_COMMIT} -X \"main.datetime=${BUILD_DATE}\" -X \"main.goVersion=$(go version)\"" \
-gcflags="-trimpath=${ROOT_DIR} -m" \
${ROOT_DIR}/main.go

echo "[Build Success]"

echo "[Build Dir] ${BUILD_DIR}"
ls -Alh ${BUILD_DIR}
