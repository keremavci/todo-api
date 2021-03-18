FROM golang:1.14.12-stretch as builder
WORKDIR $GOPATH/src/github.com/keremavci/todo-api
ADD . .
RUN go mod download && \
    go test -v ./... && \
    CGO_ENABLED=0 GOOC=linux GOARCH=amd64 go build -o todo-api .


FROM scratch
EXPOSE 8080
COPY --from=builder /go/src/github.com/keremavci/todo-api/todo-api /todo-api
ENTRYPOINT ["/todo-api"]