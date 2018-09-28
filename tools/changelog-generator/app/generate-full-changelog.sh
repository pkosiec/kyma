#!/usr/bin/env sh

source ${APP_PATH}/variables.sh
source ${APP_PATH}/config-setup.sh

FIRST_COMMIT=$(git rev-list --max-parents=0 HEAD)
echo "Getting changes starting from first commit '${FIRST_COMMIT}'..."
lerna-changelog --from=${FIRST_COMMIT} | sed -e "s/## Unreleased/## ${LATEST_VERSION}/g" > ${FULL_CHANGELOG_FILE_PATH}

source ${APP_PATH}/config-cleanup.sh