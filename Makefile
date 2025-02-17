.PHONY: all clean \
		help \
		test test-race test-coverage-report test-coveralls \
		bench bench-compare \
		upgrade \
		lint \
		sec-scan sec-trivy-scan sec-vuln-scan \
		gci-format

help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

SHELL = /bin/bash

########
# test #
########

test: ## launch all tests
	go test ./... -cover -race -failfast

test-coverage-report: ## test with coverage report
	go test -v ./... -coverpkg=./... -cover -race -covermode=atomic -coverprofile=./tmp.coverage.out && cat tmp.coverage.out | grep -v '/mocks' > ./coverage.out
	go tool cover -html=coverage.out

test-coveralls:
	go test ./... -coverpkg=./... -race -failfast -covermode=atomic -coverprofile=./tmp.coverage.out && cat tmp.coverage.out | grep -v '/mocks' > ./coverage.out
	go tool goveralls -covermode=atomic -coverprofile=./coverage.out -repotoken=$(COVERALLS_TOKEN)

test-clean-cache: ## clean test cache
	go clean -testcache


#############
# benchmark #
#############

bench: ## launch benchs
	go test ./... -bench=. -benchmem | tee ./bench.txt

bench-compare: ## compare benchs results
	go tool benchstat ./bench.txt


############
# upgrades #
############

upgrade: ## upgrade dependencies (beware, it can break everything)
	go mod tidy && \
	go get -t -u ./... && \
	go mod tidy


########
# lint #
########

lint: ## lints the entire codebase
	@go tool golangci-lint run ./... --config=./.golangci.toml

lint-clean-cache: ## clean the linter cache
	@go tool golangci-lint cache clean

lint-goleak:
	@find . -type d | while read -r dir; do \
	  if ls "$$dir"/*_test.go > /dev/null 2>&1; then \
	    if ! grep -q 'func TestMain' "$$dir"/*_test.go; then \
	      echo "Error: No TestMain found in $$dir"; \
	      exit 1; \
	    else \
	      if ! grep -q 'goleak.VerifyTestMain' "$$dir"/*_test.go; then \
	        echo "Error: goleak.VerifyTestMain missing in TestMain in $$dir"; \
	        exit 1; \
	      fi \
	    fi \
	  fi \
	done


#######
# sec #
#######

sec-scan: sec-trivy-scan sec-vuln-scan ## scan for security and vulnerability issues

sec-trivy-scan: ## scan for sec issues with trivy (trivy binary needed)
	trivy fs --exit-code 1 --no-progress --severity CRITICAL ./

sec-vuln-scan: ## scan for vulnerability issues with govulncheck (govulncheck binary needed)
	go tool govulncheck ./...

gci-format: ## format repo through gci linter
	go tool gci write ./ --skip-generated -s standard -s default -s "Prefix(github.com/induzo/enumgen)"

############
# generate #
############

generate: ## generate code
	go generate ./...
