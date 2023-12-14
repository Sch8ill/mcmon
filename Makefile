bin_name=mcmon
target=cmd/main.go

all: build

run:
	go run $(target)

build:
	go build -o build/$(bin_name) $(target)

clean:
	rm -rf build