FROM golang:1.16-buster
WORKDIR /go/src/app
COPY ./banktransfer .
RUN apt update && apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go mod download
RUN go generate ./...
RUN go install
CMD ["banktransfer"]
EXPOSE 9111