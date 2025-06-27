API_DESIGN_SRCS := $(shell find api -path api/node_modules -prune -o -print)

all: openapi.yml
	
openapi.yml: $(API_DESIGN_SRCS) api/node_modules/.bin/tsp
	cd api && tsp compile .
	
api/node_modules/.bin/tsp:
	cd api && npm i

openapi-generated/go.mod: openapi.yml
	mkdir -p openapi-generated
	cd openapi-generated && \
		npx openapi-generator-cli generate -i ../openapi.yml -g go \
	    --additional-properties=isGoSubmodule=true,useSingleRequestParameter=true,withInterfaces=true \
		sed -i 's/sdf/sdf/' openapi-generated/go.mod


openapi-generated: openapi-generated/.openapi-generator/FILES

clean:
	rm -rf openapi.yml openapi-generated