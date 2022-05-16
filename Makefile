.SILENT: #don't echo commands as we run them.

run: ## Start environment
	docker-compose up -d

stop: ## Stop environment
	docker-compose down

run-tests: ## Run unit-tests
	go test ./...

update-docs: ## Update Swagger Documentation *(Update docs before run the container).
	swag init