.PHONY: backend/gen
backend/gen: backend/gen-controller

backend/gen-controller:
	@oapi-codegen -config backend/config/oapi-codegen.yaml backend/.docs/api/openapi.yaml