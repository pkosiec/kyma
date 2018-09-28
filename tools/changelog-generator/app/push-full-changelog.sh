#!/usr/bin/env sh

SSH_FILE=$1
CONFIGURE_GIT=$2

if [$CONFIGURE_GIT == "--configure-git"]; then
    # configure git
    git config --global user.email "kyma.bot@sap.com"
    git config --global user.name "Kyma Bot"
    git config --global core.sshCommand 'ssh -i '$SSH_FILE''
fi

git add ./CHANGELOG.md
git commit -m "Update CHANGELOG.md for version ${LATEST_VERSION}"
git push origin HEAD:master
gitResult=$?
