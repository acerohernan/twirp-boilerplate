FROM golang:1.21-alpine as builder

WORKDIR /usr/app

COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# copy source
COPY pkg/ pkg/
COPY rpc/ rpc/
COPY main.go ./

RUN go build -o main main.go

FROM alpine

COPY --from=builder /usr/app/main /main

ENTRYPOINT [ "/main" ]