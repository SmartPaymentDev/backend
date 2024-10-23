BINARY=hicolleagues-crm-be

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -mod=vendor -a -installsuffix cgo -ldflags '-w'

dependencies:
	@echo "> Installing the server dependencies ..."
	@go mod download
	@go mod vendor

run:
	@echo "--- running server ---"
	@go run main.go