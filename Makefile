.PHONY: backend/gen
backend/gen:
	@docker run --rm -v "${PWD}/backend:/local" openapitools/openapi-generator-cli generate \
		-i /local/.docs/api/openapi.yaml \
		-g go-echo-server \
		-o /local \
		-p packageName=todo -p packageVersion=1.22.6
	@sh ./scripts/backend_gen.sh