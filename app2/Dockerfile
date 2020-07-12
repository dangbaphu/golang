FROM golang:1.14
WORKDIR /go/src/app2
COPY . .
RUN go get github.com/githubnemo/CompileDaemon

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT CompileDaemon --build="go build -o main" --command="./main"