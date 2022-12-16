test:
	go test -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

fmt:
	go fmt ./...

run:
	go run main.go