.PHONY: generate

# Regenerate the Go API client from the OpenAPI/Swagger spec.
# Version is pinned in openapitools.json (currently 7.14.0).
#
# --type-mappings=DateTime=Time maps every `format: date-time` field to the
# custom Time type defined in time.go instead of the stdlib time.Time. This is
# required because the Dispatch API returns datetimes both with and without a
# timezone (e.g. "2026-05-11T00:00:00"), which stdlib time.Time cannot decode.
# See the "Date formats" section of the API docs.
#
# The post-generation sed re-applies a hand patch to the generated client.go:
# the query/header serializer only special-cases stdlib time.Time, so the
# custom Time would serialize as garbage ("openapi.Time value"). Asserting the
# custom Time keeps date query parameters (e.g. dateFrom/dateTo) correct.
generate:
	npx @openapitools/openapi-generator-cli generate -i swagger.json -g go -o . --type-mappings=DateTime=Time
	sed -i '' 's/obj.(time.Time)/obj.(Time)/' client.go
