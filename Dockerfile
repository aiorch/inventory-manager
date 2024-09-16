FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o inventory-manager .

FROM scratch AS runtime
COPY --from=build /go/src/inventory-manager ./
EXPOSE 8080/tcp
ENTRYPOINT ["./inventory-manager"]
