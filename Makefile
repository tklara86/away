include .env

server:
	go run ./cmd/web

gulp:
	npm start

migrateup:
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/away?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/away?sslmode=disable" -verbose down


