.PHONY: run
run: generate-swagger
	go run main.go

.PHONY: generate-swagger
generate-swagger:
	@swag init --quiet --output openapi
