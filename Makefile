test:
	go test -v ./... -bench . -cover
run:
	go run main.go
install:
	go get .
