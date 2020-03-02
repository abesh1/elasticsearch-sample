.PHONY: gen
gen: ## Go Generate
	@which statik > /dev/null || go install github.com/rakyll/statik
	@go generate config/config.go

.PHONY: seed
seed: gen ## Insert Seed
	@go run main.go seed --env=local

.PHONY: local_build
local_build: gen ## Go Build
	@go build -o $(GOPATH)/bin/elasticsearch_sample main.go

.PHONY: build
build: gen ## Go Build
	@sh misc/scripts/ci/build.sh

.PHONY: bundle
bundle: ## Install vender library
	@go mod download
	@go mod verify

.PHONY: archive
archive: ## Archive binary file...
	@sh misc/scripts/ci/archive.sh

.PHONY: deploy
deploy: ## Deploy
	@sh misc/scripts/ci/deploy.sh