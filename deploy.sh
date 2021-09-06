#!/bin/bash

set -e

FUNCTION_NAME=go2jen
ZIP_NAME=func.zip

function init {
  yc serverless function create --name=$FUNCTION_NAME
  echo 'Remember to make function public!'
}

function createZip {
  zip -1 -R $ZIP_NAME '*.go' '*/*.go' '*/*/*.go' '*/*/*/*.go' '*.mod' '*.sum' '*/*/*.mod' '*/*/*.sum'
}

function upload {
  createZip
  yc serverless function version create \
    --function-name=go2jen \
    --runtime golang116 \
    --entrypoint main.Handler \
    --memory 128m \
    --execution-timeout 3s \
    --source-path ./func.zip
  echo 'All set!'
}

CMD=$1
shift

$CMD $@
