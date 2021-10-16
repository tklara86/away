server:
	go run ./cmd/web

gulp:
	npm start

migrateup:
	migrate -path migrations -database "postgresql://tomasz:oDX3VgHVtNQQ2lSN@165.232.111.81:5432/away?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://tomasz:oDX3VgHVtNQQ2lSN@165.232.111.81:5432/away?sslmode=disable" -verbose down


