#!/bin/bash

set -eux

function uploader_golang() {
    source /sh/setup-gruploader.sh
    source /sh/setup-golang.sh
    source /sh/release-golang.sh
}
function uploader_golangs() {
    source /sh/setup-gruploader.sh
    source /sh/setup-golang.sh
    source /sh/release-golangs.sh
}

function uploader_node() {
    source /sh/setup-gruploader.sh
    source /sh/setup-node.sh
    source /sh/release-node.sh
}

function uploader_files() {
    source /sh/setup-gruploader.sh
    source /sh/release-extra-files.sh
}

function node_gh_page() {
    source /sh/setup-node.sh
    source /sh/gh-page.sh
}

function uploader_test() {
    echo "This is just a test action output ..."
}

INPUT_ACTION="${INPUT_ACTION:-"go"}"

# 使用 switch 判断并执行对应的函数
case "${INPUT_ACTION}" in
    "golang")
        uploader_golang
        ;;
    "golangs")
        uploader_golangs
        ;;
    "node")
        uploader_node
        ;;
    "files")
        uploader_files
        ;;
    "node-gh-page")
        node-gh-page
        ;;
    "test")
        uploader_test
        ;;
    *)
        echo "Unknown action: ${INPUT_ACTION} Still working hard on development ..."
        exit 1
        ;;
esac