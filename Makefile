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

# Test Tasks
test:
	go generate ./core/mocks
	cd core && \
		go test -coverprofile=coverage.out -v ./... && \
		go tool cover -html=coverage.out -o .
	
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
	
# Development tasks
server:
	cd adapters/inbound/local && go run .
	
# Cleanup Tasks
clean-lambdas:
	rm -rf lambdas/books/bootstrap
	rm -rf lambdas/infra/artifacts
	rm -rf lambdas/books/books.zip
	rm -rf infra/plan
	rm -rf lambdas/infra/plan
	
clean: clean-lambdas
	rm -rf adapters/inbound/http/openapi.yml
	rm -rf api/node_modules
	rm -rf lambdas/gen.go