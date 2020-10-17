#!/usr/bin/env bash

set -e

rm -rf ${GCR_REPO}
git clone https://yinjiayi:${GITHUB_TOKEN}@github.com/yinjiayi/gcr.git ${GCR_REPO}
