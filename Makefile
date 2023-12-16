default: build

build:
	@go build

run: build
	@./dui-interpreter

run-input: build
	@./dui-interpreter input.dui

clean:
	@rm dui-interpreter
