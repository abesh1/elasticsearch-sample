#!/bin/bash

# import
DIR=$(cd $(dirname $0) && pwd)
. ${DIR}/base.sh

# clean
rm -Rf ${DEPLOY_DIR}
rm -Rf ${ARCHIVE_DIR}

# archive
mkdir -p ${DEPLOY_DIR} ${ARCHIVE_DIR}
ls -al ${DEPLOY_DIR}

cp ${BUILD_DIR}/${APP_BIN} ${DEPLOY_DIR}/
cp ${ROOT_DIR}/misc/scripts/codedeploy/* ${DEPLOY_DIR}/

ls -al ${DEPLOY_DIR}

# compress
tar -C ${DEPLOY_DIR} -cvzf ${ARCHIVE_DIR}/${APP_TAR_FILE} ./

ls -al ${ARCHIVE_DIR}
