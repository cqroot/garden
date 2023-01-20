PROJ_NAME=garden
BUILD_DIR=$(CURDIR)/.build

.PHONY: build
build:
	@cd $(CURDIR)/ui && npm run build
	@go build -o "$(BUILD_DIR)/$(PROJ_NAME)" $(CURDIR)/main.go

.PHONY: run
run: build
	@GARDEN_LOG_LEVEL=Debug $(BUILD_DIR)/$(PROJ_NAME)

.PHONY: dev
dev:
	@GARDEN_LOG_LEVEL=Debug nodemon -e "go" --exec go run main.go --signal SIGTERM

.PHONY: test
test:
	go test -v ./...

.PHONY: http-test
http-test:
	@bash $(CURDIR)/scripts/http-test.sh
