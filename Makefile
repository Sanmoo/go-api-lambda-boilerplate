API_DESIGN_SRCS := api/tspconfig.yaml api/main.tsp

all: openapi-generated
	
openapi: api/openapi.yml
	
api/openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

lambdas/gen.go: api/openapi.yml
	go tool oapi-codegen -config api/oid-codegen.yml api/openapi.yml

openapi-generated: lambdas/gen.go
	
clean:
	rm -rf api/openapi.yml
	rm -rf api/node_modules
	rm -rf lambdas/gen.go