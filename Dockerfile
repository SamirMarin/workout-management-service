FROM --platform=$BUILDPLATFORM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app


FROM alpine

WORKDIR /app

# add binary
COPY --from=builder /app/app app

ENTRYPOINT ["/app/app"]