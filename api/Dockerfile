FROM golang:1.19-bullseye

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

# hot-reload
RUN go install github.com/githubnemo/CompileDaemon@latest
ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main