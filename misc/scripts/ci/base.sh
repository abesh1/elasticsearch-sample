#!/bin/bash

# app
APP_BIN=elasticsearch_sample

# app
ROOT_DIR=$(cd $(dirname $0)/../../.. && pwd)
BUILD_DIR=${ROOT_DIR}/build
DEPLOY_DIR=${BUILD_DIR}/deploy
ARCHIVE_DIR=${BUILD_DIR}/archive

# git
SHORT_COMMIT=$(git rev-parse --short HEAD)
LONG_COMMIT=$(git rev-parse --verify HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
PATH_BRANCH=$(echo $BRANCH | sed -e "s/\//-/")

# date
LONG_DATE=$(date +%Y%m%d%H%M%S)
SHORT_DATE=$(date +%Y%m%d)

# build
TAR_FILE_SUFFIX=${PATH_BRANCH}-${SHORT_COMMIT}.tar.gz
APP_TAR_FILE=${APP_BIN}-${TAR_FILE_SUFFIX}

# deploy
S3_PATH_SUFFIX=${PATH_BRANCH}/${SHORT_DATE}/${LONG_DATE}-${SHORT_COMMIT}

# s3 bucket
S3_BUCKET=elasticsarch-sample-deploy
