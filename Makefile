all: build

build: build-poodr-ch2

build-poodr-ch2:
	go build -o ./bin/poodr-ch2 ./cmd/poodr-ch2

run-poodr-ch2:
	go run ./cmd/poodr-ch2

run-poodr-ch3:
	go run ./cmd/poodr-ch3

test:
	go test ./...

clean:
	rm ./bin/poodr-ch2
	rmdir bin