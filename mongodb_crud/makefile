SHELL:=/bin/sh

run:
	docker-compose up --build

generate:
	protoc --go_out=plugins=grpc:. todopb/todo.proto

clean:
	rm -fr todopb/todo.pb.go

fmt:
.PHONY: run generate clean
