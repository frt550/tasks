# Tasks

How to run:
1. `go mod install` to download modules
2. `buf generate` to generate code for grpc/rest api by .proto files 
3. `docker-compose up` to start db
4. ./migrate.sh to apply migrations
5. `make run-tg` to run telegram bot api
   - or `make run-task` to run rest api for task
   - or `make run-backup` to run rest api for backup

Other commands:
- `make .deps` to generate protoc binaries
- `make migration NAME=your_name_here` to create migration with specified name
- `./migrate.sh` to apply new migrations

Some info:
- Telegram bot username `fb_TasksBot`