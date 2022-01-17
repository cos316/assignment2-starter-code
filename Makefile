all: client server

clean:
	rm -f client server *.exe

client: $(wildcard cmd/microblog-client/*.go)
	go build -o client ./cmd/microblog-client

server: $(wildcard cmd/microblog-server/*.go) $(wildcard http_router/*.go)
	go build -o server ./cmd/microblog-server
