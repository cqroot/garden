PROJ_NAME=todo
BUILD_DIR=$(CURDIR)/.build

.PHONY: build
build:
	go build -o "$(BUILD_DIR)/$(PROJ_NAME)" main.go

.PHONY: run
run: build
	@$(BUILD_DIR)/$(PROJ_NAME)

.PHONY: test
test:
	go test -v ./...

.PHONY: http-test
http-test:
	@bash $(CURDIR)/scripts/http-test.sh
