clean:
	@rm -rf build

build: clean
	@go build -o build/cmng main.go

run-new: build
	./build/cmng new