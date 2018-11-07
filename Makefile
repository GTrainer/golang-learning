export GOPATH = ./

all: learning

learning:
	go build -o $@ ./src/

clean: learning
	rm -rf $<

