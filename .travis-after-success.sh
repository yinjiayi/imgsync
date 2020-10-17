#!/usr/bin/env bash

set -e

export TZ=UTC-8
cd ${GCR_REPO}
if [ -n "$(git status --porcelain)" ]; then
  git add .
  git commit -m "Travis CI Auto Update(`date +'%Y-%m-%d %H:%M:%S'`)"
  git push https://yinjiayi:${GITHUB_TOKEN}@github.com/yinjiayi/gcr.git
fi
