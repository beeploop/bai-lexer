default: build

build:
	@go build

run: build
	@./dui-interpreter

clean:
	@rm dui-interpreter
