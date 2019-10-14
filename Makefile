.PHONY: build

wiregen:
	wire cmd/wire.go

mockgen:
	mockgen -source repository/entry.go -destination mock_repository/mock_entry.go
	mockgen -source repository/user.go -destination mock_repository/mock_user.go

build: wiregen
	go build -o bin/sample cmd/main.go cmd/wire_gen.go

test: mockgen test/service

test/service:
	go test service/*
