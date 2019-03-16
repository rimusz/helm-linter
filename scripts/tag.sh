#!/bin/sh -e

# shellcheck disable=SC2002
tag="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Tagging helm-linter with ${tag} ..."

git checkout master
git tag -a -m "Release $tag" "$tag" 
git push origin "$tag"
