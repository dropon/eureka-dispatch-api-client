.PHONY: generate

# Regenerate the Go API client from the OpenAPI/Swagger spec.
# Version is pinned in openapitools.json (currently 7.14.0).
generate:
	npx @openapitools/openapi-generator-cli generate -i swagger.json -g go -o .
