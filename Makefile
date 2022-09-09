all:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/rds-auth-token.linux-amd64 ./cmd/rds-auth-token
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/rds-auth-token.darwin-amd64 ./cmd/rds-auth-token
