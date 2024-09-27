.PHONY: backend/gen
backend/gen: backend/gen-controller backend/gen-models

backend/gen-controller:
	@oapi-codegen -config backend/config/oapi-codegen.yaml backend/.docs/api/openapi.yaml

backend/gen-models:
	@sqlboiler psql -o backend/internal/repository/models

.PHONY: install-tools
install-tools:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.0
	@go install github.com/volatiletech/sqlboiler/v4@latest
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	@go install go.uber.org/mock/mockgen@v0.4.0

.PHONY: migrate-create
migrate-create:
	@read -p "Please enter name: " input;\
	migrate create -dir backend/migrations -seq -ext sql $$input

.PHONY: migrate-up
migrate-up:
	@migrate -source file://backend/migrations -database postgres://dev:dev@localhost:5432/dev?sslmode=disable up 1

.PHONY: migrate-down
migrate-down:
	@migrate -source file://backend/migrations -database postgres://dev:dev@localhost:5432/dev?sslmode=disable down 1

.PHONY: migrate-force
migrate-force:
	@migrate -source file://backend/migrations -database postgres://dev:dev@localhost:5432/dev?sslmode=disable force 1

.PHONY: migrate-reset
migrate-reset: migrate-drop migrate-up
	
.PHONY: migrate-drop
migrate-drop:
	@migrate -source file://backend/migrations -database postgres://dev:dev@localhost:5432/dev?sslmode=disable drop