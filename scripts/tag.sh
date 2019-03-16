#!/bin/sh -e

# shellcheck disable=SC2002
tag="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Tagging helm-linter with v${tag} ..."

git checkout master
git pull
git tag -a -m "Release v$tag" "v$tag" 
git push origin "v$tag"
