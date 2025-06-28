API_DESIGN_SRCS := api/tspconfig.yaml api/main.tsp

# General tasks
all: openapi-generated lambdas
	
openapi: api/openapi.yml
	
api/openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

adaptors/inbound/aws/gen.go: api/openapi.yml
	go tool oapi-codegen -config api/oid-codegen.yml api/openapi.yml

openapi-generated: adaptors/inbound/http/gen.go
	
# Lambda Build tasks
lambdas: books

books: adaptors/inbound/aws/infra/artifacts/books.zip

adaptors/inbound/aws/infra/artifacts/books.zip: adaptors/inbound/aws/books/bootstrap
	
adaptors/inbound/aws/books/bootstrap: adaptors/inbound/aws/books/lambda.go
	cd adaptors/inbound/aws/books && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap lambda.go
	cd adaptors/inbound/aws/books && zip books.zip bootstrap
	mkdir -p adaptors/inbound/aws/infra/artifacts
	mv adaptors/inbound/aws/books/books.zip adaptors/inbound/aws/infra/artifacts
	
# Deployment tasks
terraform-init:
	cd infra && terraform init
	cd adaptors/inbound/aws/infra && terraform init

plan-infra-deployment:
	cd infra && terraform plan -out=plan

deploy-infra:
	cd infra && terraform apply plan

plan: lambdas
	cd adaptors/inbound/aws/infra && terraform plan -out=plan
	
make plan-destroy:
	cd adaptors/inbound/aws/infra && terraform plan -destroy -out=plan

deploy:
	cd adaptors/inbound/aws/infra && terraform apply plan
	
# Development tasks
server:
	go run adaptors/inbound/local/main.go
	
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