# helm-linter plugin

Helm plugin to find in chart's `values.yaml` a key (case sensitive) which contains word password and check the key with it isn’t empty.

## Install

Based on the version in `plugin.yaml`, release binary will be downloaded from GitHub:

```
$ helm plugin install https://github.com/rimusz/helm-linter
Downloading and installing helm-linter v0.1.0 ...
https://github.com/rimusz/helm-linter/releases/download/v0.1.0/helm-linter_0.1.0_darwin_amd64.tar.gz
Installed plugin: linter
```
