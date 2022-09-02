ARG port=8078

FROM golang as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/grpc ./cmd/task/grpc/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/rest ./cmd/task/rest/main.go

# final stage
FROM frolvlad/alpine-bash
ARG port
COPY --from=builder /out/grpc /grpc
COPY --from=builder /out/rest /rest
COPY .env /.env

EXPOSE ${port}
COPY docker/service/wrapper-script.sh wrapper-script.sh
RUN chmod +x wrapper-script.sh
CMD ./wrapper-script.sh