############################
# STEP 1 build executable binary
############################
FROM golang:1.18-alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.s
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/github.com/samlitowitz/go-say
RUN pwd
COPY . .
# Fetch dependencies.
# Using go mod.
ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-say cmd/go-say/*.go

# Test
FROM alpine:3.16 AS test

COPY --from=builder /go/bin/go-say /usr/local/bin/go-say
COPY --from=builder /go/src/github.com/samlitowitz/go-say/assets/fixtures /tmp/test/expected
COPY --from=builder /go/src/github.com/samlitowitz/go-say/assets/templates /usr/local/share/cows
COPY --from=builder /go/src/github.com/samlitowitz/go-say/scripts/test.sh /app/test.sh
