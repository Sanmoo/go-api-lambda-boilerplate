module github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/aws

go 1.24.4

require (
	github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http v0.0.0-00010101000000-000000000000
	github.com/akrylysov/algnhsa v1.1.0
)

require (
	github.com/Sanmoo/go-api-lambda-boilerplate/core v0.0.0-20250628120149-1dc9e5f4ee3b // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/aws/aws-lambda-go v1.47.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
)

replace github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http => ../http
