# helm-linter plugin

Helm plugin to find in chart's `values.yaml` a key which contains word password and check the key isn’t empty.

## Install

Based on the version in `plugin.yaml`, release binary will be downloaded from GitHub:

```
$ helm plugin install https://github.com/rimusz/helm-linter
Downloading and installing helm-linter v0.1.1 ...
https://github.com/rimusz/helm-linter/releases/download/v0.1.1/helm-linter_0.1.1_darwin_amd64.tar.gz
Installed plugin: linter
```

## Usage

Check charts for hard-coded passwords:

```
$ helm linter testdata/mychart/
Checking chart testdata/mychart/

Found hard-coded password/s:
 mysql.aaa.bPassword: eafaefawf
 mysql.aaa.cPassword.password: dfafasf
 mysql.mysqlRootPassword: druidroot
 mysql.mysqlPassword: druid
 password: sadasdsad
```

```
$ helm linter testdata/postgresql/
Checking chart testdata/postgresql/

Hard-coded password/s have not been found :-)
```
