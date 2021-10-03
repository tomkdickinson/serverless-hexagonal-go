#!/bin/bash
set -eu

if [ -f ./go.mod ]; then
    exit 0
fi

if [ $# -eq 0 ]; then
  PROJECT_NAME=$(basename $(pwd | xargs dirname))
  CURRENT_DIR=$(basename $(pwd))
  MODULE_NAME=github.com/${PROJECT_NAME}/${CURRENT_DIR}
else
  MODULE_NAME=$1
fi

echo $MODULE_NAME
touch go.mod

CONTENT=$(cat <<-EOD
module ${MODULE_NAME}

require (
  github.com/aws/aws-lambda-go v1.6.0
  github.com/google/wire v0.5.0
  github.com/rs/zerolog v1.25.0
)
EOD
)

echo "$CONTENT" > go.mod
grep -rl github.com/tomkdickinson/serverless-hexagonal-go --exclude={gomod.sh,README.md} . | xargs sed -i "s:github.com/tomkdickinson/serverless-hexagonal-go:$MODULE_NAME:g"
go mod tidy