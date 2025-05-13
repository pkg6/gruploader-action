#!/bin/bash -eux

INPUT_CNAME_DOMAIN="${INPUT_CNAME_DOMAIN:-""}"
INPUT_ORIGIN_URL="${INPUT_ORIGIN_URL:-""}"
INPUT_NPM_INSTALL_PKG="${INPUT_NPM_INSTALL_PKG:-"npm ci --legacy-peer-deps"}"
INPUT_NPM_BUILD_COMMAND="${INPUT_NPM_BUILD_COMMAND:-"npm run build"}"
INPUT_NPM_BUILD_DIST="${INPUT_NPM_BUILD_DIST:-"dist"}"

if [ -z "$INPUT_ORIGIN_URL" ]; then
  INPUT_ORIGIN_URL=$(git config --get remote.origin.url || true)
fi

if [ -z "$INPUT_ORIGIN_URL" ]; then
  if [ -n "$GITHUB_REPOSITORY" ]; then
    INPUT_ORIGIN_URL="https://github.com/${GITHUB_REPOSITORY}.git"
    echo "Using GitHub Actions repository URL: $INPUT_ORIGIN_URL"
  else
    echo "Error: Could not determine origin URL from INPUT_ORIGIN_URL, git config, or GITHUB_REPOSITORY." >&2
    exit 1
  fi
fi


echo "Step 1: Installing dependencies..."
${INPUT_NPM_INSTALL_PKG} || { echo "install pkg failed"; exit 1; }

echo "Step 2: Building project with command:  ${INPUT_NPM_BUILD_COMMAND}..."
${INPUT_NPM_BUILD_COMMAND} || { echo "npm run ${NPM_RUN_BUILD_CMD} failed"; exit 1; }


if [ ! -d "$INPUT_NPM_BUILD_DIST" ]; then
  echo "path $INPUT_NPM_BUILD_DIST not exit" >&2
  exit 1
fi

if [ -n "$INPUT_CNAME_DOMAIN" ]; then
  echo "Custom CNAME domain detected: $INPUT_CNAME_DOMAIN"
  echo "$INPUT_CNAME_DOMAIN" > "$INPUT_NPM_BUILD_DIST/CNAME" || { echo "Failed to write CNAME"; exit 1; }
elif [ -f "CNAME" ]; then
  echo "Found local CNAME file, copying to $INPUT_NPM_BUILD_DIST..."
  cp "CNAME" "$INPUT_NPM_BUILD_DIST/" || { echo "Failed to copy CNAME file"; exit 1; }
fi

echo "Step 3: Entering build directory..."
cd ${INPUT_NPM_BUILD_DIST} || { echo "cd ${INPUT_NPM_BUILD_DIST} failed"; exit 1; }


echo "Step 4: Preparing to publish..."

git init

git config user.name "github-actions[bot]"
git config user.email "41898282+github-actions[bot]@users.noreply.github.com"

git remote add origin "$INPUT_ORIGIN_URL" || { echo "git remote add origin failed"; exit 1; }
git branch -M gh-pages  || { echo "git branch -M gh-pages failed"; exit 1; }
git add -A || { echo "git add -A failed"; exit 1; }
git commit -m "$(date "+gruploader-action push %Y-%m-%d %H:%M:%S")" || { echo "git commit failed"; exit 1; }
git config -l
git push -u origin gh-pages --force || { echo  "Git push failed"; exit 1; } 