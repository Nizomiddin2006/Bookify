SRC_DIR = src
BIN_DIR = bin

TARGET = $(BIN_DIR)/bookify

SRCS = $(wildcard $(SRC_DIR)/*.go)

all: $(TARGET)

$(TARGET): $(SRCS)
	mkdir -p $(BIN_DIR)
	go build -o $@ $^

clean:
	rm -rf $(BIN_DIR)

.PHONY: all clean
