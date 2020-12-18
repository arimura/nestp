run:
	go run cmd/main.go

run-sample:
	cat data/sample.json | $(MAKE) run