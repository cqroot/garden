.PHONY: run-server
run-server:
	@go run $(CURDIR)/cmd/todo-server/main.go

.PHONY: http-test
http-test:
	@bash $(CURDIR)/scripts/http-test.sh
