# Disable automatic deletion of intermediary files
.SECONDARY:

API_DESIGN_SRCS := api/tspconfig.yaml api/main.tsp

# General tasks
all: openapi-generated lambdas
	
openapi: adapters/inbound/http/openapi.yml
	
adapters/inbound/http/openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

adapters/inbound/http/gen.go: adapters/inbound/http/openapi.yml api/oapi-codegen.yml
	go tool oapi-codegen -config api/oapi-codegen.yml adapters/inbound/http/openapi.yml

openapi-generated: adapters/inbound/http/gen.go
	
# Lambda Build tasks
LAMBDA_NAMES := books movies tv-series non-electronic-games electronic-games
lambdas: $(addprefix build-,$(LAMBDA_NAMES))

build-%: adapters/inbound/aws/infra/artifacts/%.zip
	@echo "Built lambda: $*"

adapters/inbound/aws/infra/artifacts/%.zip: adapters/inbound/aws/%/bootstrap
	cd adapters/inbound/aws/$* && zip $*.zip bootstrap
	mkdir -p adapters/inbound/aws/infra/artifacts
	mv adapters/inbound/aws/$*/$*.zip adapters/inbound/aws/infra/artifacts
	
HTTP_SRCS := $(shell find adapters/inbound/http -name "*.go")
AWS_COMMON_SRCS := $(shell find adapters/inbound/aws -name "*.go" ! -name "lambda.go")

adapters/inbound/aws/%/bootstrap: adapters/inbound/aws/%/lambda.go $(HTTP_SRCS) $(AWS_COMMON_SRCS)
	cd adapters/inbound/aws/$* && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap lambda.go

# Deployment tasks
terraform-init:
	cd infra && terraform init
	cd adapters/inbound/aws/infra && terraform init

plan-infra:
	cd infra && terraform plan -out=plan

deploy-infra:
	cd infra && terraform apply plan

plan:
	cd adapters/inbound/aws/infra && terraform plan -out=plan
	
make plan-destroy:
	cd adapters/inbound/aws/infra && terraform plan -destroy -out=plan

deploy:
	cd adapters/inbound/aws/infra && terraform apply plan
	
# Development and Testing tasks
CORE_SRCS := $(shell find core -name "*.go")

.PHONY: server
server:
	cd adapters/inbound/local && go run .

.PHONY: deps
deps:
	go install go.uber.org/mock/mockgen@latest
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: test
test: core/cover.out adapters/inbound/http/cover.out

core/cover.out: $(CORE_SRCS)
	go generate ./core/mocks
	cd core && go test ./... -coverprofile=./cover.out.tmp -covermode=atomic -coverpkg=./...
	cd core && cat ./cover.out.tmp | grep -v "_generated.go" > ./cover.out

adapters/inbound/http/cover.out: $(HTTP_SRCS)
	cd adapters/inbound/http && \
		go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	
.PHONY: check-coverage
check-coverage: core/cover.out adapters/inbound/http/cover.out
	cd core && go tool cover -html=cover.out -o cover.html
	cd adapters/inbound/http && go tool cover -html=cover.out -o cover.html
	go-test-coverage --config=./.testcoverage.yaml
	
# Tasks for CI Server
.PHONY: check-coverage-ci
check-coverage-ci: core/cover.out
	go-test-coverage --config=./.testcoverage.yaml
	
# Cleanup Tasks
.PHONY: clean-lambdas
clean-lambdas:
	rm -rf lambdas/books/bootstrap
	rm -rf lambdas/infra/artifacts
	rm -rf lambdas/books/books.zip
	rm -rf infra/plan
	rm -rf lambdas/infra/plan
	
.PHONY: clean
clean: clean-lambdas
	rm -rf adapters/inbound/http/openapi.yml
	rm -rf api/node_modules
	rm -rf lambdas/gen.go