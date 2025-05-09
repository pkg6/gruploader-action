# !/bin/bash -eux

INPUT_GRUPLOADER_VERSION="${INPUT_GRUPLOADER_VERSION:-"v0.1.0"}"

INPUT_GITHUB_TOKEN="${INPUT_GITHUB_TOKEN:-}"
#Forcefully covering
INPUT_OVERWRITE="${INPUT_OVERWRITE:-true}"
INPUT_GRUPLOADER_CLOUD="${INPUT_GRUPLOADER_CLOUD:-"github"}"
#Repositories（pkg6/gruploader）
INPUT_REPOSITORY="${INPUT_REPOSITORY:-${GITHUB_REPOSITORY}}"
#Published tags
INPUT_TAG="${INPUT_TAG:-${GITHUB_REF_NAME}}"
INPUT_RETRY="${INPUT_RETRY:-3}"

debug_log() {
  echo -e "\033[32m[DEBUG] $1\033[0m"
}
error_log() {
  echo -e "\033[31m[ERROR] $1\033[0m"
}

check_empty() {
  local var_name="$1"
  local var_value="$2"
  if [ -z "${var_value}" ]; then
    error_log "${var_name} is empty. Exiting..."
    exit 1
  fi
}
check_empty "INPUT_GRUPLOADER_VERSION" "${INPUT_GRUPLOADER_VERSION}"
check_empty "INPUT_REPOSITORY" "${INPUT_REPOSITORY}"
check_empty "INPUT_TAG" "${INPUT_TAG}"
check_empty "INPUT_GRUPLOADER_CLOUD" "${INPUT_GRUPLOADER_CLOUD}"
check_empty "INPUT_OVERWRITE" "${INPUT_OVERWRITE}"
check_empty "INPUT_RETRY" "${INPUT_RETRY}"


TEMP="$(mktemp -d)"
trap 'rm -rf $TEMP' EXIT INT
wget --progress=dot:mega https://go.dev/dl/go1.23.9.linux-$(dpkg --print-architecture).tar.gz -O "$TEMP/go-linux.tar.gz"
(
    cd "$TEMP" || exit 1
    tar -zxf go-linux.tar.gz
    export GO_HOME=$TEMP/go
    export GOPATH=$TEMP/gb
    export PATH=${GOPATH}/bin:${GO_HOME}/bin/:$PATH
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
    go install github.com/pkg6/gruploader-action/gruploader@latest
    mv $TEMP/gb/bin/gruploader /usr/local/bin/gruploader
)

chmod +x /usr/local/bin/gruploader

gruploader --help

export GITHUB_TOKEN=${INPUT_GITHUB_TOKEN}
export TAG_VERSION=${INPUT_TAG}
export EXTRA_FILES=${INPUT_EXTRA_FILES}
export MAIN_GO=${INPUT_MAINGO}
export BIN_NAME=${INPUT_BIN_NAME}
export GRUPLOADER_COMMAND="gruploader -cloud ${INPUT_GRUPLOADER_CLOUD} --repo=${INPUT_REPOSITORY} --tag=${INPUT_TAG} --overwrite=${INPUT_OVERWRITE} --retry=${INPUT_RETRY}"