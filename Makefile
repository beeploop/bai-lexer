default: build

build:
	@go build -o bai

run: build
	@./bai

run-input: build
	@./bai input.bai

clean:
	@rm bai*
