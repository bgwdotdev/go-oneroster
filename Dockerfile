FROM golang:latest as builder
WORKDIR /go/src/github.com/fffnite/go-oneroster
ADD . /go/src/github.com/fffnite/go-oneroster
WORKDIR ./cmd/goors
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/goors

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/goors /
EXPOSE 80
CMD ["/goors"]
