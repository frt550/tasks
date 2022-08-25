# Tasks

How to run app:
1. `go mod download` to download modules
2. `make .deps` to generate protoc binaries
3. `buf generate` to generate code for grpc/rest api by .proto files 
4. `docker-compose up` to start db
5. `./migrate.sh` to apply migrations
6. `make run-tg` to run telegram bot api
   - or `make run-task` to run rest api for task
   - or `make run-backup` to run rest api for backup

Before running tests:
1. `go generate ./...` to generate mocks
2. `./migrate.sh test` to apply migrations on test db

Other commands:
- `make migration NAME=your_name_here` to create migration with specified name

Some info:
- Telegram bot username `fb_TasksBot`