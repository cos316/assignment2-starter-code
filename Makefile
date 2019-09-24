all: client server

clean:
	rm -f client server *.exe

client: $(wildcard microblog-client/*.go)  $(wildcard http_router/*.go)
	go build -o client ./microblog-client

server: $(wildcard microblog-server/*.go) $(wildcard http_router/*.go)
	go build -o server ./microblog-server
