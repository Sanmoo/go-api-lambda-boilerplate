API_DESIGN_SRCS := $(shell find api -path api/node_modules -prune -o -print)

all: openapi
	
openapi: api/openapi.yml
	
api/openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

lambdas/gen.go: api/openapi.yml
	go tool oapi-codegen -config ../api/oid-codegen.yml ../api/openapi.yml

openapi-generated: lambdas/gen.go
	
clean:
	rm -rf api/openapi.yml
	rm -rf api/node_modules
	rm -rf lambdas/gen.go