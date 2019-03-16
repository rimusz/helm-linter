# helm-linter plugin

Helm plugin to find in chart's `values.yaml` a key (case sensitive) which contains word password and check the key isn’t empty.

## Install

Based on the version in `plugin.yaml`, release binary will be downloaded from GitHub:

```
$ helm plugin install https://github.com/rimusz/helm-linter
Downloading and installing helm-linter v0.0.8 ...
https://github.com/rimusz/helm-linter/releases/download/v0.0.8/helm-linter_0.0.8_darwin_amd64.tar.gz
Installed plugin: linter
```

## Usage

Check the chart for hard-coded passwords:

```
helm linter testdata/mychart/values.yaml
mysql.mysqlRootPassword: druidroot
mysql.mysqlPassword: druid
mysql.aaa.bPassword: eafaefawf
mysql.aaa.cPassword.password: dfafasf
password: sadasdsad
```

