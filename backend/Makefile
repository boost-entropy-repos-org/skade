run:
	go run cmd/skade/main.go

createdb:
	docker exec -it skadedb createdb --username=skadeuser --owner=skadeuser skadedb

postgres:
	docker run -d --name skadedb -p 5432:5432 -e POSTGRES_USER=skadeuser -e POSTGRES_PASSWORD=test postgres:12.3

test:
	go test -v -cover ./...

migrateup:
	 migrate -path db/migrations -database "postgresql://skadeuser:skadepw@127.0.0.1:5432/skadedb?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migrations -database "postgresql://skadeuser:skadepw@127.0.0.1:5432/skadedb?sslmode=disable" -verbose down
