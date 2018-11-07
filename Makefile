export GOPATH = /home/centos/golang-learning

all: learning

learning:
	go build -o $@ ./src/

clean:
	rm -rf learning

