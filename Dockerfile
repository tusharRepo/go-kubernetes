FROM golang:1.14 as builder
WORKDIR /workspace
COPY ./webservice/ . 
RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix cgo -o webservice

FROM scratch
WORKDIR /
COPY --from=builder /workspace/webservice .
EXPOSE 8080
ENTRYPOINT ["/webservice"]