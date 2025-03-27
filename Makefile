build:
	@go build -o dist/service-user-runserver ./cmd/runserver
	@go build -o dist/service-user-migrate ./cmd/migrate

migrate:
	@./dist/service-user-migrate

run:
	@./dist/service-user-runserver
