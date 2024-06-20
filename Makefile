TARGET = picsum-cli

build:
	@go build -o bin/$(TARGET)

install:
	install bin/$(TARGET) /usr/local/bin