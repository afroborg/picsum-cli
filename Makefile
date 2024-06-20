TARGET = picsum-cli

build:
	@go build -o bin/$(TARGET)

install: build
	install bin/$(TARGET) /usr/local/bin