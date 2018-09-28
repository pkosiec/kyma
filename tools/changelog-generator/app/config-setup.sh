#!/usr/bin/env sh

echo "Copying configuration file..."
cp ${APP_PATH}/package.json ./

echo "Changing git settings..."
git config --global user.email "kyma.bot@sap.com"
git config --global user.name "Kyma Bot"
git config --global core.sshCommand 'ssh -i '$SSH_FILE''

echo "Temporary removing 'latest' tag..."
LATEST_TAG_REV=$(git rev-list -n 1 ${LATEST_TAG})
LATEST_TAG_MESSAGE=$(git tag -l --format='%(contents)' ${LATEST_TAG})
git tag -d ${LATEST_TAG}
