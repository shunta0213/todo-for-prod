.PHONY: backend/gen
backend/gen:
	docker run --rm -v "${PWD}/backend:/local" openapitools/openapi-generator-cli generate \
		-i /local/.docs/api/openapi.yaml \
		-g go-echo-server \
		-o /local