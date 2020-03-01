.PHONY: gen
gen: ## Go Generate
	@which statik > /dev/null || go install github.com/rakyll/statik
	@go generate config/config.go

.PHONY: seed
seed: gen ## Insert Seed
	@go run main.go seed --env=local

.PHONY: build
build: gen ## Go Build
	@go build -o $(GOPATH)/bin/elasticsearch_sample main.go