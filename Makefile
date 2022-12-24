.PHONY: generate-mocks
generate-mocks:
	mockgen -package=bots -source=./v2/internal/reqhandler/reqhandler.go > ./v2/internal/mocks/mock_reqhandler.go
