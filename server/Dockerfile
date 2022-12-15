FROM golang:1.18-alpine AS builder
ENV GOPATH=/root/go
RUN mkdir workdir
COPY  . /workdir
WORKDIR /workdir
# Install any required modules
RUN go mod download
# Build binaries
RUN go build -o bin/apiserver main.go

############# Multi stage build ##################
FROM alpine
COPY --from=builder /workdir/bin/ /
EXPOSE 8080
ENTRYPOINT ["/apiserver"]
