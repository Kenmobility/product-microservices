proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

postgres: 
	docker run --name product_microservice_con -p 5438:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it product_microservice_con createdb --username=root --owner=root product_microservice_db

dropdb:
	docker exec -it product_microservice_con dropdb product_microservice_db

  .PHONY: proto postgres createdb dropdb