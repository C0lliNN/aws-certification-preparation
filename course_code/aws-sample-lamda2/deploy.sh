#! /usr/bin/bash

version=$(make upload-code | jq -r '.VersionId')
echo $version
make version="$version" update-stack