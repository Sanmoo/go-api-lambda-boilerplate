API_DESIGN_SRCS := api/tspconfig.yaml api/main.tsp

# General tasks
all: openapi-generated lambdas
	
openapi: api/openapi.yml
	
api/openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

adapters/inbound/http/gen.go: api/openapi.yml api/oapi-codegen.yml
	go tool oapi-codegen -config api/oapi-codegen.yml api/openapi.yml

openapi-generated: adapters/inbound/http/gen.go
	
# Lambda Build tasks
lambdas: books

books: adapters/inbound/aws/infra/artifacts/books.zip

adapters/inbound/aws/infra/artifacts/books.zip: adapters/inbound/aws/books/bootstrap
	
adapters/inbound/aws/books/bootstrap: adapters/inbound/aws/books/lambda.go
	cd adapters/inbound/aws/books && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap lambda.go
	cd adapters/inbound/aws/books && zip books.zip bootstrap
	mkdir -p adapters/inbound/aws/infra/artifacts
	mv adapters/inbound/aws/books/books.zip adapters/inbound/aws/infra/artifacts
	
# Deployment tasks
terraform-init:
	cd infra && terraform init
	cd adapters/inbound/aws/infra && terraform init

plan-infra-deployment:
	cd infra && terraform plan -out=plan

deploy-infra:
	cd infra && terraform apply plan

plan: lambdas
	cd adapters/inbound/aws/infra && terraform plan -out=plan
	
make plan-destroy:
	cd adapters/inbound/aws/infra && terraform plan -destroy -out=plan

deploy:
	cd adapters/inbound/aws/infra && terraform apply plan
	
# Development tasks
server:
	go run adapters/inbound/local/main.go
	
# Cleanup Tasks
clean-lambdas:
	rm -rf lambdas/books/bootstrap
	rm -rf lambdas/infra/artifacts
	rm -rf lambdas/books/books.zip
	rm -rf infra/plan
	rm -rf lambdas/infra/plan
	
clean: clean-lambdas
	rm -rf api/openapi.yml
	rm -rf api/node_modules
	rm -rf lambdas/gen.go