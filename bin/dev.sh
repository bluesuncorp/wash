#!/bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"

echo "Script Running From $DIR"

ROOT=$DIR/..

cd $ROOT
PWD="$(pwd)"

echo "PWD=$PWD"
EXECUTABLE="$(basename $PWD)"

echo "Executable = $EXECUTABLE"
echo "Setup ENV vars"

export APP_DOMAIN=localhost
export APP_IS_PRODUCTION=false
export APP_APP_PORT=3005
export APP_SMTP_SERVER=localhost
export APP_SMTP_USERNAME=
export APP_SMTP_PASSWORD=
export APP_SMTP_PORT=1025
export APP_SUPPORT_EMAIL=support@email.com

justdoit -watch="./" -include="(.+\.go|.+\.c|.+\.yaml)$" -build="go install -v" -run="$GOPATH/bin/$EXECUTABLE"