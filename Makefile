PROJ_NAME=garden
BUILD_DIR=$(CURDIR)/.build

.PHONY: build-ui
build-ui:
	@cd $(CURDIR)/ui && npm run build

.PHONY: build-go
build-go:
	@cd $(CURDIR)/internal/app && wire
	@go build -o "$(BUILD_DIR)/$(PROJ_NAME)" $(CURDIR)/main.go

.PHONY: build
build: build-ui build-go

.PHONY: run-without-ui
run-without-ui: build-go
	@cd $(CURDIR)/internal/app && wire
	@GARDEN_LOG_LEVEL=Debug $(BUILD_DIR)/$(PROJ_NAME)

.PHONY: run
run: build-ui run-without-ui

.PHONY: clean
clean:
	rm -rf "$$(go run main.go print-data-path)"

.PHONY: ui-dev
ui-dev:
	@cd $(CURDIR)/ui && npm run dev

.PHONY: test
test:
	go test -v ./...

.PHONY: http-test
http-test:
	@bash $(CURDIR)/scripts/http-test.sh
