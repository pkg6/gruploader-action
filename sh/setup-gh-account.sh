#!/bin/bash

: "${GITHUB_ACTOR:?is not set}"
: "${INPUT_GITHUB_TOKEN:?is not set}"

echo "machine github.com login $GITHUB_ACTOR password $INPUT_GITHUB_TOKEN" > ~/.netrc

chmod 600 ~/.netrc
