default: build

build:
	@go build -o bai

build-linux:
	@go build -o releases/linux/bai

build-win:
	@env GOOS=windows GOARCH=amd64 go build -o releases/windows/bai.exe

run: build
	@./bai

run-input: build
	@./bai input.bai

clean:
	@rm bai*
