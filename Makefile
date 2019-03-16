PLUGIN_NAME := linter

HAS_DEP := $(shell command -v dep;)
HAS_PIP := $(shell command -v pip;)
HAS_VENV := $(shell command -v virtualenv;)

.PHONY: build
build: build_linux build_mac build_windows

build_windows: export GOARCH=amd64
build_windows: export GO111MODULE=on
build_windows: export GOPROXY=https://gocenter.io
build_windows:
	@GOOS=windows go build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/windows/amd64/helmlinter helmlinter/main.go  # windows

link_windows:
	@cp bin/windows/amd64/helmlinter ./bin/helmlinter

build_linux: export GOARCH=amd64
build_linux: export CGO_ENABLED=0
build_linux: export GO111MODULE=on
build_linux: export GOPROXY=https://gocenter.io
build_linux:
	@GOOS=linux go build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/linux/amd64/helmlinter helmlinter/main.go  # linux

link_linux:
	@cp bin/linux/amd64/helmlinter ./bin/helmlinter

build_mac: export GOARCH=amd64
build_mac: export CGO_ENABLED=0
build_mac: export GO111MODULE=on
build_mac: export GOPROXY=https://gocenter.io
build_mac:
	@GOOS=darwin go build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/darwin/amd64/helmlinter helmlinter/main.go # mac osx
	@cp bin/darwin/amd64/helmlinter ./bin/helmlinter # For use w make install

link_mac:
	@cp bin/darwin/amd64/helmlinter ./bin/helmlinter

.PHONY: tag
tag:
	@scripts/tag.sh

.PHONY: clean
clean:
	@git status --ignored --short | grep '^!! ' | sed 's/!! //' | xargs rm -rf

.PHONY: install
install:
	HELM_LOGIN_PLUGIN_NO_INSTALL_HOOK=1 helm plugin install $(shell pwd)

.PHONY: remove
remove:
	helm plugin remove $(PLUGIN_NAME)
