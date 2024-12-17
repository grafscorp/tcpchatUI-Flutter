all: clean server client

clean:
	rm -rf ./build/*

server: 
	go build -o ./build/server ./src/server/main.go ./src/server/global_objects.go ./src/server/handle_connection.go ./src/server/listen_cli.go ./src/server/listen_connection.go ./src/server/sender.go ./src/server/logger.go ./src/server/parse_args.go

client:
	go build -o ./build/client ./src/client/main.go ./src/client/parse_args.go

run_server:
	./build/server port :50051 pwd 123

run_client:
	./build/client address 127.0.0.1:50051 name DefaultName pwd 123