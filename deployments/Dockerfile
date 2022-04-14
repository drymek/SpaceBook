### Builder
FROM golang:1.18-alpine as builder
RUN apk add --no-cache openssh git build-base
RUN go install github.com/cespare/reflex@latest
WORKDIR /go/src/sb
COPY go.mod /go/src/sb
RUN go mod download

COPY . /go/src/sb/

RUN go build -o sb /go/src/sb/cmd/app

### Production
FROM alpine:3.15
ENV PORT 8080
EXPOSE 8080

RUN apk add --no-cache curl
RUN addgroup -S sb && adduser -S -G sb sb

COPY --chown=sb --from=builder /go/src/sb/sb /go/app/

USER sb

HEALTHCHECK --interval=5m --timeout=3s \
   CMD curl -f http://localhost:${PORT}/heathcheck || exit 1

CMD ["/go/app/sb"]