build:
	docker-compose up --build -d api

start:
	docker-compose up -d

migration-up:
	migrate -path ./schema -database 'postgres://postgres:pass@localhost:54320/postgres?sslmode=disable' up