.PHONY: all clean

all: gobooru

gobooru: *.go */*.go
	go build

clean:
	rm -rf gobooru
