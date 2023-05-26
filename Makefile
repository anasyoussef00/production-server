BINARY_NAME=main.out

all: build run

build:
	go build -o ${BINARY_NAME} cmd/main.go

test:
	go test -v cmd/main.go

run:
	go build -o ${BINARY_NAME} cmd/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}