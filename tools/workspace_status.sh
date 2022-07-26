#!/usr/bin/env bash

set -euo pipefail

commit_sha=$(git rev-parse HEAD)
echo "STABLE_GIT_COMMIT $commit_sha"

git_branch=$(git rev-parse --abbrev-ref HEAD)
echo "STABLE_GIT_BRANCH $git_branch"

git_tree_status=$(git diff-index --quiet HEAD -- && echo 'Clean' || echo 'Modified')
echo "GIT_TREE_STATUS $git_tree_status"

#version=$(cat VERSION)
#echo "VERSION $version"
