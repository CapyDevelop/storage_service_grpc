gen_python:
	python -m grpc_tools.protoc -I. --python_out=./storage --grpc_python_out=./storage storage_service.proto

gen_go:
	protoc --go_out=. --go-grpc_out=. storage_service.proto