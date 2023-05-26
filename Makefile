BINARY_NAME=main.out

all: clean build run

build:
	go build -o ${BINARY_NAME} cmd/main.go

test:
	go test -v cmd/main.go

run:
	./${BINARY_NAME}

clean:
	go clean