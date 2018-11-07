export GOPATH = ./

all: learning

learning:
	go build -o $@ ./src/

clean:
	rm -rf learning

