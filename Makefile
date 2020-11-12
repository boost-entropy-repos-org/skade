run:
	go run cmd/skade/main.go

createdb:
	docker exec -it skadedb createdb --username=skadeuser --owner=skadeuser skadedb

postgres:
	docker run -d --name skadedb -p 5432:5432 -e POSTGRES_USER=skadeuser -e POSTGRES_PASSWORD=secret postgres:12.3


