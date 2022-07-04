BINARY = server
FILE = main.go

main: $(FILE)
	go mod download
	go build -o $(BINARY) $(FILE)

clear:
	sudo rm -rf ./uploads/**

test: $(FILE)
	go test
