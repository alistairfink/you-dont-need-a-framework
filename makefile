.PHONY: run
run:
	go run cmd/main.go

.PHONY: run-jq
run-jq:
	go run cmd/main.go | jq .