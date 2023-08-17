.PHONY: clean

clean:
	rm -rf bin

build: clean
	mkdir -p bin
	go build -o bin/daily-rover-api-go ./main.go

run:
	./bin/daily-rover-api-go
	

fmt:
	go fmt ./... && go vet ./...
	
all: fmt clean build run