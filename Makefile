all: build

build: build-poodr-ch2

build-poodr-ch2:
	go build -o bin/poodr-ch2 ./cmd/poodr-ch2

run-poodr-ch2:
	go run ./cmd/poodr-ch2

clean:
	rm bin/poodr-ch2
	rmdir bin