#!/bin/bash

# import
DIR=$(cd $(dirname $0) && pwd)
. ${DIR}/base.sh

# https://docs.aws.amazon.com/cli/latest/reference/deploy/create-deployment.html

S3_KEY=${APP_BIN}/${S3_PATH_SUFFIX}/${APP_BIN}.tar.gz

# upload api
aws s3 cp ${ARCHIVE_DIR}/${APP_TAR_FILE} \
s3://${S3_BUCKET}/${S3_KEY}

# deploy api
aws deploy create-deployment \
--application-name ${APP_BIN} \
--deployment-group-name Dev \
--deployment-config-name CodeDeployDefault.AllAtOnce \
--file-exists-behavior OVERWRITE \
--s3-location bucket=${S3_BUCKET},key=${S3_KEY},bundleType=tgz
